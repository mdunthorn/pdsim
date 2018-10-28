#!/usr/bin/python

from binascii import hexlify
import getopt
import logging
import select
import socket
import sys
import time


def main():
    cf = setup()
    logging.info("start program")
    logging.debug(cf)

    host = cf["host"]
    port = cf["port"]

    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    sock.connect((host, port))

    req = "\xff\x8f\x01\x01\x01"
    ba = bytearray()
    ready_to_send = True
    last_sent_time = time.time()

    while True:

        input_ready, output_ready, error_ready = select.select([sock], [sock], [])

        if (0 < len(input_ready)):
            chunk = sock.recv(2048)
            logging.debug("read a {} byte chunk".format(len(chunk)))
            ba.extend(chunk)
            message = None
            ba, messages = decode(ba)
            if messages is not None:
                for m in messages:
                    logging.info("decoded a {} message".format(m))

        if ready_to_send and (0 < len(output_ready)):
            sent = sock.send(req)
            logging.info("sent {} bytes: {}".format(sent, hexlify(req)))
            last_sent_time = time.time()
            ready_to_send = False

        now = time.time()
        if cf["poll_cycle"] <= (now - last_sent_time):
            ready_to_send = True


def decode(ba):
    logging.debug("decoding {}".format(hexlify(ba)))
    messages = []
    while True:
        logging.debug("try to decode {}".format(hexlify(ba)))
        if 0 == len(ba):
            break

        if 0xff != ba[0]:
            ba = find_start_of_message(ba)
    
        if 3 <= len(ba):
            message_type = ba[1]
            logging.debug("found message_type {}".format(hex(message_type)))
            message_len = ba[2]
    
            need_bytes = 3 + message_len + 1
    
            if need_bytes == len(ba):
                messages.append(hex(message_type))
                ba = bytearray()
            elif need_bytes < len(ba):
                messages.append(hex(message_type))
                ba = ba[need_bytes:]
        else:
            break

    return ba, messages


def find_start_of_message(ba):
    i = 0
    while i < len(ba):
        if 0xff == ba[i]:
            return ba[i:]
    return bytearray()


def setup():
    opts = parse_command_line()
    setup_logging(opts["log_level"])
    return opts


def parse_command_line():
    opts, args = getopt.getopt(sys.argv[1:],
        "h:p:P:l:H",
        ["host=", "port=", "poll_cycle=", "log_level=", "help"])

    opt_dict = {
        "log_level": "info",
        "poll_cycle": 20
    }

    try:
        for o, a in opts:
            if o in ("-H", "--help"):
                usage()
                sys.exit(0)
            elif o in ("-h", "--host"):
                opt_dict["host"] = a
            elif o in ("-p", "--port"):
                opt_dict["port"] = int(a)
            elif o in ("-l", "--log_level"):
                opt_dict["log_level"] = a
            elif o in ("-P", "--poll_cycle"):
                opt_dict["poll_cycle"] = int(a)
    except getopt.GetoptError, e:
        print "error parsing command line: ", e
        usage()
        sys.exit(1)

    if (not "host" in opt_dict) or (not "port" in opt_dict):
        usage()
        sys.exit(1)
        
    opt_dict["log_level"] = get_log_level(opt_dict["log_level"])

    return opt_dict


def get_log_level(s):
    log_level = logging.INFO
    s = s.upper()
    if "DEBUG" == s:
        log_level = logging.DEBUG
    elif "INFO" == s:
        log_level = logging.INFO
    elif "WARNING" == s:
        log_level = logging.WARNING
    elif "ERROR" == s:
        log_level = logging.ERROR
    elif "CRITICAL" == s:
        log_level = logging.CRITICAL
    return log_level


def setup_logging(log_level):
    log_format = "%(asctime)-15s %(levelname)-8s %(module)s - %(message)s"
    logging.basicConfig(level=log_level, format=log_format)
    return True


def usage():
    print """usage: client.py <options>
  -h, --host        detector hostname or ip address
  -i, --poll_cycle  poll interval in seconds
  -p, --port        detector port
  -l, --log_level   log level (info, debug, error, warning)
  -H, --help        print this message
"""


if __name__ == "__main__":
    main()

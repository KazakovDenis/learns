#!/usr/bin/env python3
import os
import sys
import redis


producer = redis.Redis()
file = os.path.basename(__file__)


def main():
    if len(sys.argv) < 3:
        raise Exception(f'Usage: python {file} [channel] [message]')

    channel, message = sys.argv[1], sys.argv[2:]
    active = producer.pubsub_channels()

    if channel.encode() not in active:
        print('No active channels to send message to.')
        sys.exit(1)

    producer.publish(channel, message)


if __name__ == '__main__':
    main()

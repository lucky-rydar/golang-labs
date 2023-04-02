#! /usr/bin/python3

import argparse
import requests
import json

SERVER_ADDR = 'http://localhost:8080'
TOKEN_FILE = 'token'

def parser_init():
    parser = argparse.ArgumentParser(description='Dormitory Management System')

    parser.add_argument('cmd', type=str, help='subcommand to execute')
    
    # user operations
    parser.add_argument('-u', '--username', type=str, help='username')
    parser.add_argument('-p', '--password', type=str, help='password')


    # ask operations
    parser.add_argument('--name', type=str, help='name')
    parser.add_argument('--surname', type=str, help='surname')
    parser.add_argument('--is_male', action='store_true', help='is male')
    parser.add_argument('--ticket_number', type=str, help='ticket number')
    parser.add_argument('--ticket_expire_date', type=str, help='ticket expire date, format=YYYY-MM-DD')
    parser.add_argument('--room_number', type=str, help='room number')

    # resolve action
    parser.add_argument('--action_id', type=int, help='action id')
    parser.add_argument('--is_approved', type=bool, help='is approved action', default=True)

    # add room
    parser.add_argument('--area_sqm', type=int, help='area sqm')
    

    return parser

def parser_args():
    return parser_init().parse_args()


def get_token():
    try:
        with open(TOKEN_FILE, 'r') as f:
            return f.read()
    except Exception as e:
        print(e)
        print('Please login first')
        exit(1)


def http_req(url, json):
    try:
        response = requests.get(url, json=json)
        if response.status_code != 200:
            print(f'Request failed {response.status_code} {response.text}')

        return response
    except Exception as e:
        print(e)
        print('Request failed')
        return


def register(args):
    username = args.username
    password = args.password

    data = {
        'username': username,
        'password': password
    }

    url = SERVER_ADDR + '/user/register'

    response = http_req(url, data)

    if response.status_code == 200:
        print('Register successfully')


def login(args):
    username = args.username
    password = args.password

    data = {
        'username': username,
        'password': password
    }

    url = SERVER_ADDR + '/user/login'

    response = http_req(url, data)

    if response.status_code == 200:
        json_data = json.loads(response.text)
        token = json_data['Uuid']

        # save token to local file
        with open(TOKEN_FILE, 'w') as f:
            f.write(token)

        print('Login successfully')


def ask_register(args):
    name = args.name
    surname = args.surname
    is_male = args.is_male
    ticket_number = args.ticket_number
    ticket_expire_date = args.ticket_expire_date
    
    data = {
        'name': name,
        'surname': surname,
        'isMale': is_male,
        'studentTicketNumber': ticket_number,
        'studentTicketExpireDate': ticket_expire_date+'T00:00:00Z'
    }

    url = SERVER_ADDR + '/ask_admin/register'

    response = http_req(url, data)

    if response.status_code == 200:
        print('Ask register successfully')

        print(f'{response.text}')


def ask_contract_sign(args):
    ticket_number = args.ticket_number

    data = {
        'studentTicketNumber': ticket_number
    }

    url = SERVER_ADDR + '/ask_admin/contract/sign'

    response = http_req(url, data)

    if response.status_code == 200:
        print('Ask contract sign successfully')


def ask_settle(args):
    ticket_number = args.ticket_number
    room_number = args.room_number

    data = {
        'studentTicketNumber': ticket_number,
        'roomNumber': room_number
    }

    url = SERVER_ADDR + '/ask_admin/settle'

    response = http_req(url, data)

    if response.status_code == 200:
        print('Ask settle successfully')


def ask_unsettle(args):
    ticket_number = args.ticket_number

    data = {
        'studentTicketNumber': ticket_number
    }

    url = SERVER_ADDR + '/ask_admin/unsettle'

    response = http_req(url, data)

    if response.status_code == 200:
        print('Ask unsettle successfully')


def ask_resettle(args):
    ticket_number = args.ticket_number
    room_number = args.room_number

    data = {
        'studentTicketNumber': ticket_number,
        'roomNumber': room_number
    }

    url = SERVER_ADDR + '/ask_admin/resettle'

    response = http_req(url, data)

    if response.status_code == 200:
        print('Ask resettle successfully')


def ask_actions():
    token = get_token()

    data = {
        'uuid': token
    }

    url = SERVER_ADDR + '/ask_admin/actions'

    response = http_req(url, data)

    if response.status_code == 200:
        # parse response
        actions = json.loads(response.text)
        if len(actions) == 0:
            print('No actions')
            return

        for action in actions:
            print(json.dumps(action, indent=4))


def resolve_action(args):
    is_approved = args.is_approved
    action_id = args.action_id
    token = get_token()

    data = {
        'uuid': token,
        'is_approved': is_approved,
        'action_id': action_id
    }

    url = SERVER_ADDR + '/ask_admin/actions/resolve'

    response = http_req(url, data)

    if response.status_code == 200:
        print('Action resolved')


def add_room(args):
    is_male = args.is_male
    room_number = args.room_number
    area_sqm = args.area_sqm
    uuid = get_token()

    data = {
        'number': room_number,
        'isMale': is_male,
        'areaSqMeters': area_sqm,
        'uuid': uuid
    }

    print(json.dumps(data, indent=4))

    url = SERVER_ADDR + '/rooms/add'

    response = http_req(url, data)

    if response.status_code == 200:
        print('Room added')


def get_dorm_stats():
    token = get_token()

    data = {
        'uuid': token
    }

    url = SERVER_ADDR + '/dormitory/load/stats'

    response = http_req(url, data)

    if response.status_code == 200:
        print(json.dumps(json.loads(response.text), indent=4))


def get_students():
    token = get_token()

    data = {
        'uuid': token
    }

    url = SERVER_ADDR + '/students'

    response = http_req(url, data)

    if response.status_code == 200:
        print(json.dumps(json.loads(response.text), indent=4))

def main():
    args = parser_args()

    if args.cmd == 'register':
        register(args)
    elif args.cmd == 'login':
        login(args)
    elif args.cmd == 'ask_register':
        ask_register(args)
    elif args.cmd == 'ask_contract_sign':
        ask_contract_sign(args)
    elif args.cmd == 'ask_settle':
        ask_settle(args)
    elif args.cmd == 'ask_unsettle':
        ask_unsettle(args)
    elif args.cmd == 'ask_resettle':
        ask_resettle(args)
    elif args.cmd == 'ask_actions':
        ask_actions()
    elif args.cmd == 'resolve':
        resolve_action(args)
    elif args.cmd == 'add_room':
        add_room(args)
    elif args.cmd == 'get_dorm_stats':
        get_dorm_stats()
    elif args.cmd == 'students':
        get_students()
    else:
        print('Unknown command')

if __name__ == '__main__':
    main()

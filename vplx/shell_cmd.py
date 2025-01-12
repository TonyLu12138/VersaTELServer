#!/usr/bin/env python
'''
This file is part of ConfigShell.
Copyright (c) 2011-2013 by Datera, Inc

Licensed under the Apache License, Version 2.0 (the "License"); you may
not use this file except in compliance with the License. You may obtain
a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
License for the specific language governing permissions and limitations
under the License.
'''

import os
from iscsi_json import JsonOperation
from configshell_fb import ConfigNode,ConfigShell, ExecutionError
import os

class MySystemRoot(ConfigNode):
    def __init__(self, shell):
        ConfigNode.__init__(self, '/', shell=shell)
        Interpreters(self)
        System(self)
        for level in range(1,5):
            ConfigNode("level%d" % level, self)

class Interpreters(ConfigNode):
    def __init__(self, parent):
        ConfigNode.__init__(self, 'interpreters', parent)

    def ui_command_python(self):
        '''
        python  - an interpreted object-oriented programming language
        '''
        os.system("python")

    def ui_command_ipython(self):
        '''
        ipython - An Enhanced Interactive Python
        '''
        os.system("ipython")

    def ui_command_bash(self):
        '''
        bash - GNU Bourne-Again SHell
        '''
        os.system("bash")

class System(ConfigNode):
    def __init__(self, parent):
        ConfigNode.__init__(self, 'system', parent)
        Processes(self)
        Storage(self)

    def ui_command_uname(self):
        '''
        Displays the system uname information.
        '''
        os.system("uname -a")

    def ui_command_lsmod(self):
        '''
        lsmod - program to show the status of modules in the Linux Kernel
        '''
        os.system("lsmod")

    def ui_command_lspci(self):
        '''
        lspci - list all PCI devices
        '''
        os.system("lspci")

    def ui_command_lsusb(self):
        '''
        lsusb - list USB devices
        '''
        os.system("lsusb")

    def ui_command_lscpu(self):
        '''
        lscpu - CPU architecture information helper
        '''
        os.system("lscpu")

    def ui_command_uptime(self):
        '''
        uptime - Tell how long the system has been running.
        '''
        os.system("uptime")


class Storage(ConfigNode):
    def __init__(self, parent):
        ConfigNode.__init__(self, 'storage', parent)

    def ui_command_lsscsi(self):
        '''
        lsscsi - list SCSI devices (or hosts) and their attributes
        '''
        os.system("lsscsi")

    def ui_command_du(self):
        '''
        du - estimate file space usage
        '''
        os.system("du -hs /")

    def ui_command_df(self):
        '''
        df - report file system disk space usage
        '''
        os.system("df -h")

    def ui_command_uuidgen(self):
        '''
        uuidgen - command-line utility to create a new UUID value.
        '''
        os.system("uuidgen")

class Processes(ConfigNode):
    def __init__(self, parent):
        ConfigNode.__init__(self, 'processes', parent)

    def ui_command_top(self):
        '''
        top - display Linux tasks
        '''
        os.system("top")

    def ui_command_ps(self):
        '''
        ps - report a snapshot of the current processes.
        '''
        os.system("ps aux")

    def ui_command_pstree(self):
        '''
        pstree - display a tree of processes
        '''
        os.system("pstree")

class Users(ConfigNode):
    def __init__(self, parent):
        ConfigNode.__init__(self, 'users', parent)

    def ui_command_who(self):
        '''
        who - show who is logged on
        '''
        os.system("who")

    def ui_command_whoami(self):
        '''
        whoami - print effective userid
        '''
        os.system("whoami")

    def ui_command_users(self):
        '''
        users  -  print the user names of users currently logged in.
        '''
        os.system("users")

def main():
    shell = ConfigShell('~/.myshell')
    root_node = MySystemRoot(shell)
    shell.run_interactive()

if __name__ == "__main__":
    main()

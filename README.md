## Technical task

1. Set-Up a WireGuard VPN server. Once a user is connected to the server, he should be able to use a DNS server on the VPN host (that is only accessible to VPN-connected clients). This DNS server should forward queries to root DNS servers (it should not use a resolver outside of the VPN server), however, all queries to  netflix.com and its subdomains should be rewritten and routed to a locally running sniproxy service. This sniproxy should then route traffic out to actual Netflix. Ensure Netflix web page can load for a client via sniproxy.
2. Secure the VPN server to your understanding
3. Automate your set-up using a configuration tool of your choice (SaltStack, Ansible, Chef, Puppet, etc). Ensure to use best practices according to your chosen tool
4. Provide a working WireGuard client config (so we can connect and verify all is working as intended), as well as your automation code.
5. Write a NordVPN public API parser in your chosen language (Python, Ruby, etc.) that says which NordVPN servers are in which subnet.
	- Parser should parse a configuration file that includes the API url ([https://api.nordvpn.com/v1/servers](https://api.nordvpn.com/v1/servers)), a subnet, and a prefix size that the subnet should be split into  
	- Parser should accept command line arguments for a config file, log level, log file
	- It should split the subnet into multiple subnets of the given prefix size  
	- It should pull the list of NordVPN servers
	- It should then group servers into generated subnets and print this information out  
	- Optionally the script should support multiple log levels  
	- Use best practices of your chosen language when writing the script
	- Below you will find an example of how running your parser might look like.

```bash
# cat /etc/nordsec/config.ini
api_url=https://api.nordvpn.com/v1/servers
subnet=192.168.0.0/24
new_prefix=25
# ./task.py --config /etc/nordsec/config.ini --log_level=debug --log=/var/log/nordsec.log
Subnet: 192.168.0.0/25 Hosts: [ tst1.nordvpn.com ]
Subnet: 192.168.0.128/25 Hosts: [ tst2.nordvpn.com ]
Hosts with no subnet match: [ tst3.nordvpn.com ]
# cat /var/log/nordsec.log
DEBUG: Loading Configuration: /etc/nordsec/config.ini
DEBUG: Loaded subnet: 192.168.0.0/24
DEBUG: Generating new prefix size: 25
INFO: Generated subnets: [ 192.168.0.0/25, 192.168.0.128/25 ]
DEBUG: Pulling info from https://api.nordvpn.com/v1/servers
DEBUG: Loaded 10 servers from API
DEBUG: Filtered 3 online servers from API
DEBUG: Host: tst1.nordvpn.com IP: 192.168.0.15 Subnet: 192.168.0.0/25 Status: match
DEBUG: Host: tst1.nordvpn.com IP: 192.168.0.15 Subnet: 192.168.0.128/25 Status: miss
DEBUG: Host: tst2.nordvpn.com IP: 192.168.0.160 Subnet: 192.168.0.0/25 Status: miss
DEBUG: Host: tst2.nordvpn.com IP: 192.168.0.160 Subnet: 192.168.0.128/25 Status: match
DEBUG: Host: tst3.nordvpn.com IP: 192.168.1.160 Subnet: 192.168.0.0/25 Status: miss
DEBUG: Host: tst3.nordvpn.com IP: 192.168.1.160 Subnet: 192.168.0.128/25 Status: miss
INFO: Subnet: 192.168.0.0/25 Hosts: [ tst1.nordvpn.com ]
INFO: Subnet: 192.168.0.128/25 Hosts: [ tst2.nordvpn.com ]
INFO: Hosts with no subnet match: [ tst3.nordvpn.com ]
```

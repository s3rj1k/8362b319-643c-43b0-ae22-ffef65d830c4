# salt-ssh

Cheetsheet for using 'salt-ssh'

## Presetup

- Change current working directory in terminal into this directory so that 'Saltfile' will be in the CWD.
- Update './config/roster' file with proper SSH key location for each roster record (key 'priv')

### Test connection
`salt-ssh -i '*' test.ping` or `salt-ssh \* test.ping` or `salt-ssh vpn test.ping`

### Test SLS config setup
`salt-ssh vpn state.apply ping_test test=True`

### Debug state.apply
`salt-ssh vpn state.apply ping_test test=True -l debug` or `salt-ssh vpn state.apply ping_test test=True`

### Get pillar items
`salt-ssh '*' pillar.items`

### Test highstate
`salt-ssh vpn state.highstate test=True`

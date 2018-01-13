# Ryzen Stabilizator Tabajara

Simple Go program to disable C6 C-state on an AMD Ryzen processor, in order to help with the infamous "MCE-random-reboots-while-idle" issue.

Code licensed under Apache License 2.0.

## Usage:

### Check status of C6 C-state:
```
sudo ./ryzen-stabilizator
C6 C-state is DISABLED.
```

### Enable C6 C-state:
```
sudo ./ryzen-stabilizator --enable-c6
Enabling C6 C-state:   SUCCESS
```

### Disable C6 C-state:
```
sudo ./ryzen-stabilizator --disable-c6
Disabling C6 C-state:   SUCCESS
```


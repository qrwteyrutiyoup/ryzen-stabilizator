# Ryzen Stabilizator Tabajara

Simple Go program to enable/disable C6 C-state, processor boosting and address space layout randoization (ASLR) on an AMD Ryzen processor, in order to help with the infamous "MCE-random-reboots-while-idle" issue.

Code licensed under Apache License 2.0.

## Basic usage:

### Check status of C6 C-state, processor boosting and ASLR:
```
sudo ./ryzen-stabilizator
Ryzen Stabilizator Tabajara unspecified/git version
Copyright (C) 2018 Sergio Correia <sergio@correia.cc>


C6 C-state is DISABLED.
ASLR is DISABLED.
Processor boosting is ENABLED.
```

### Enable C6 C-state:
```
sudo ./ryzen-stabilizator --enable-c6
Ryzen Stabilizator Tabajara unspecified/git version
Copyright (C) 2018 Sergio Correia <sergio@correia.cc>

Enabling C6 C-state:   SUCCESS

C6 C-state is ENABLED.
ASLR is DISABLED.
Processor boosting is ENABLED.
```

### Disable C6 C-state:
```
sudo ./ryzen-stabilizator --disable-c6
Ryzen Stabilizator Tabajara unspecified/git version
Copyright (C) 2018 Sergio Correia <sergio@correia.cc>

Disabling C6 C-state:   SUCCESS

C6 C-state is DISABLED.
ASLR is DISABLED.
Processor boosting is ENABLED.
```

### Enable processor boosting:
```
sudo ./ryzen-stabilizator --enable-boosting
Ryzen Stabilizator Tabajara unspecified/git version
Copyright (C) 2018 Sergio Correia <sergio@correia.cc>

Enabling processor boosting:   SUCCESS

C6 C-state is DISABLED.
ASLR is DISABLED.
Processor boosting is ENABLED.
```

### Disable processor boosting:
```
sudo ./ryzen-stabilizator --disable-boosting
Ryzen Stabilizator Tabajara unspecified/git version
Copyright (C) 2018 Sergio Correia <sergio@correia.cc>

Disabling processor boosting:   SUCCESS

C6 C-state is DISABLED.
ASLR is DISABLED.
Processor boosting is DISABLED.
```

### Enable address space layout randomization (ASLR):
```
sudo ./ryzen-stabilizator --enable-aslr
Ryzen Stabilizator Tabajara unspecified/git version
Copyright (C) 2018 Sergio Correia <sergio@correia.cc>

Enabling address space layout randomization (ASLR):   SUCCESS

C6 C-state is DISABLED.
ASLR is ENABLED.
Processor boosting is DISABLED.
```

### Disable address space layout randomization (ASLR):
```
sudo ./ryzen-stabilizator --disable-aslr
Ryzen Stabilizator Tabajara unspecified/git version
Copyright (C) 2018 Sergio Correia <sergio@correia.cc>

Disabling address space layout randomization (ASLR):   SUCCESS

C6 C-state is DISABLED.
ASLR is DISABLED.
Processor boosting is DISABLED.
```

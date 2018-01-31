# Ryzen Stabilizator Tabajara

Simple Go program to disable C6 C-state and/or processor boosting on an AMD Ryzen processor, in order to help with the infamous "MCE-random-reboots-while-idle" issue.

Code licensed under Apache License 2.0.

## Basic usage:

### Check status of both C6 C-state and processor boosting:
```
sudo ./ryzen-stabilizator
Ryzen Stabilizator Tabajara unspecified/git version
Copyright (C) 2018 Sergio Correia <sergio@correia.cc>


C6 C-state is ENABLED.
Processor boosting is DISABLED.
```

### Enable C6 C-state:
```
sudo ./ryzen-stabilizator --enable-c6
Ryzen Stabilizator Tabajara unspecified/git version
Copyright (C) 2018 Sergio Correia <sergio@correia.cc>

Enabling C6 C-state:   SUCCESS

C6 C-state is ENABLED.
Processor boosting is DISABLED.
```

### Disable C6 C-state:
```
sudo ./ryzen-stabilizator --disable-c6
Ryzen Stabilizator Tabajara unspecified/git version
Copyright (C) 2018 Sergio Correia <sergio@correia.cc>

Disabling C6 C-state:   SUCCESS

C6 C-state is DISABLED.
Processor boosting is DISABLED.
```

### Enable processor boosting:
```
sudo ./ryzen-stabilizator --enable-boosting
Ryzen Stabilizator Tabajara unspecified/git version
Copyright (C) 2018 Sergio Correia <sergio@correia.cc>

Enabling processor boosting:   SUCCESS

C6 C-state is DISABLED.
Processor boosting is ENABLED.
```

### Disable processor boosting:
```
sudo ./ryzen-stabilizator --disable-boosting
Ryzen Stabilizator Tabajara unspecified/git version
Copyright (C) 2018 Sergio Correia <sergio@correia.cc>

Enabling processor boosting:   SUCCESS

C6 C-state is DISABLED.
Processor boosting is ENABLED.
```

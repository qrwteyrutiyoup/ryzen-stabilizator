#!/bin/bash
#
# About: Install Ryzen Stabilizator automatically
# Author: liberodark
# License: GNU GPLv3

version="0.0.1"

echo "Welcome on Ryzen Stabilizator Install Script $version"

#=================================================
# CHECK ROOT
#=================================================

if [[ $(id -u) -ne 0 ]] ; then echo "Please run as root" ; exit 1 ; fi

#=================================================
# RETRIEVE ARGUMENTS FROM THE MANIFEST AND VAR
#=================================================

distribution=$(cat /etc/*release | grep "PRETTY_NAME" | sed 's/PRETTY_NAME=//g' | sed 's/["]//g' | awk '{print $1}')

install_ryzen_stabilizator(){
      git clone https://github.com/qrwteyrutiyoup/ryzen-stabilizator.git
      pushd ryzen-stabilizator/ || exit
      export GOPATH=./go
      go get
      go build
      mv ryzen-stabilizator.conf /etc/modules-load.d/ryzen-stabilizator.conf
      chmod 0644 /etc/modules-load.d/ryzen-stabilizator.conf
      mkdir /usr/share/licenses/ryzen_stabilizator/
      mv LICENSE /usr/share/licenses/ryzen_stabilizator/LICENSE
      mv ./contrib/systemd/ryzen* /usr/lib/systemd/system/
      mkdir /etc/ryzen-stabilizator/
      mv ./contrib/settings* /etc/ryzen-stabilizator/settings.toml
      chmod 0644 /etc/ryzen-stabilizator/settings.toml
      mv ./ryzen-stabilizator-master /usr/bin/ryzen-stabilizator
      chmod 0755 /usr/bin/ryzen-stabilizator
      popd || exit
}

deps_install(){
echo "Install Ryzen Stabilizator ($distribution)"

  # Check OS & git / go

  if ! command -v ryzen-stabilizator; then

    if [[ "$distribution" = CentOS || "$distribution" = CentOS || "$distribution" = Red\ Hat || "$distribution" = Fedora || "$distribution" = Suse || "$distribution" = Oracle ]]; then
      yum install -y git go

      install_ryzen_stabilizator || exit
    
    elif [[ "$distribution" = Debian || "$distribution" = Ubuntu || "$distribution" = Deepin ]]; then
      apt-get update
      apt-get install -y git go --force-yes
    
      install_ryzen_stabilizator || exit
      
    elif [[ "$distribution" = Manjaro || "$distribution" = Arch\ Linux ]]; then
      pacman -S git go --noconfirm
    
      install_ryzen_stabilizator || exit

    fi
fi
}

deps_install

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
      git clone https://github.com/qrwteyrutiyoup/ryzen-stabilizator.git &> /dev/null
      pushd ryzen-stabilizator/ || exit
      #export GOPATH=./go
      go get &> /dev/null
      go build &> /dev/null
      touch ryzen-stabilizator.conf
      echo -e "# msr module is required by ryzen-stabilizator." >> "ryzen-stabilizator.conf"
      echo -e "msr" >> "ryzen-stabilizator.conf"
      cp -a ./ryzen-stabilizator.conf /etc/modules-load.d/ryzen-stabilizator.conf
      chmod 0644 /etc/modules-load.d/ryzen-stabilizator.conf
      mkdir -p /usr/share/licenses/ryzen_stabilizator/
      cp -a LICENSE /usr/share/licenses/ryzen_stabilizator/LICENSE
      cp -a ./contrib/systemd/ryzen* /etc/systemd/system/
      mkdir -p /etc/ryzen-stabilizator/
      cp -a ./contrib/settings* /etc/ryzen-stabilizator/settings.toml
      chmod 0644 /etc/ryzen-stabilizator/settings.toml
      cp -a ./ryzen-stabilizator /usr/bin/ryzen-stabilizator
      chmod 0755 /usr/bin/ryzen-stabilizator
      popd || exit
      echo "Install Ryzen Stabilizator Finish"
}

deps_install(){
echo "Install Ryzen Stabilizator ($distribution)"

  # Check OS & git / go

  if ! command -v ryzen-stabilizator; then

    if [[ "$distribution" = CentOS || "$distribution" = CentOS || "$distribution" = Red\ Hat || "$distribution" = Fedora || "$distribution" = Suse || "$distribution" = Oracle ]]; then
      yum install -y git golang-go &> /dev/null

      install_ryzen_stabilizator || exit
    
    elif [[ "$distribution" = Debian || "$distribution" = Ubuntu || "$distribution" = Deepin || "$distribution" = KDE ]]; then
      apt-get update
      apt-get install -y git golang-go --force-yes &> /dev/null
    
      install_ryzen_stabilizator || exit
      
    elif [[ "$distribution" = Manjaro || "$distribution" = Arch\ Linux ]]; then
      pacman -S git go --noconfirm &> /dev/null
    
      install_ryzen_stabilizator || exit

    fi
fi
}

deps_install

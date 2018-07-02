

1. 编译hyperscan为动态库形势， 并安装， 
2. 在/etc/profile后增加： export PKG_CONFIG_PATH=/usr/share/pkgconfig:/usr/lib/pkgconfig:/usr/local/lib64/pkgconfig
3. ln -s /usr/bin/pkg-config /usr/local/bin/pkg-config

TODO:
引用的不是vendor下的库， 而且$GOPATH下的库 ？？？

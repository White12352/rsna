#!/bin/bash

sudo apt-get install fuse -y

cp -r linux64 nekoray.AppDir

# The file for Appimage

rm nekoray.AppDir/launcher

cat >nekoray.AppDir/nekoray.desktop <<-EOF
[Desktop Entry]
Name=NekoXray
Exec=echo "NekoXray started"
Icon=nekoray
Type=Application
Categories=Network
EOF

cat >nekoray.AppDir/AppRun <<-EOF
#!/bin/bash
echo "PATH: \${PATH}"
echo "NekoXray runing on: \$APPDIR"
if [[ \${NKR_APPIMAGE_TUN} == "1" ]] || [[ \${NKR_APPIMAGE_TUN} == "true" ]]; then
    permissionCmd=\$(type -p pkexec || type -p gksudo || type -p sudo)
fi

export LD_LIBRARY_PATH=\${APPDIR}/usr/lib
export QT_PLUGIN_PATH=\${APPDIR}/usr/plugins
\${permissionCmd} \${APPDIR}/nekoray -appdata \$@
EOF

chmod +x nekoray.AppDir/AppRun

# build

curl -fLSO https://github.com/AppImage/AppImageKit/releases/latest/download/appimagetool-x86_64.AppImage
chmod +x appimagetool-x86_64.AppImage
./appimagetool-x86_64.AppImage nekoray.AppDir

# clean

rm appimagetool-x86_64.AppImage
rm -rf nekoray.AppDir

#!/bin/bash

wails build -platform windows/amd64 -clean

tar -czf exporter-windows-amd64.tar.gz ./ui/build/bin/safetyculture-exporter.exe

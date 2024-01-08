#!/bin/bash

# 判断配置文件是否存在
CONFIG_FILE="/app/config.yaml"
if [[ ! -e "$CONFIG_FILE" ]]; then
  echo "$CONFIG_FILE doesn't exists!"
  return
fi

/app/sso -c "$CONFIG_FILE"
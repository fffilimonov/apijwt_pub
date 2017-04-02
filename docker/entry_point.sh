#!/bin/bash

./apijwt
NODE_PID=$!
wait $NODE_PID

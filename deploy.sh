#!/bin/sh
scp -i ~/.ssh/terraform.pem app.bz2 ec2-user@107.22.153.92:/tmp/
scp -i ~/.ssh/terraform.pem index.html ec2-user@107.22.153.92:/tmp/

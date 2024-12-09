# astra-vdi-activity-analyzer

Тема 04. Захват и анализ пользовательской активности в режиме реального времени в инфраструктуре виртуальных рабочих мест (VDI) (Python-библиотека yolov5)

![Scheme](/demos/scheme.png)

## Project structure

1. `agent` - a program installed on the end-users to capture the screenshots and send to storage;
2. `storage` - a server that receives screenshots, validates agents and sends to processing service;
3. `processing` - a server that receives screenshots, processes them with machine learning and stores the data to database;
4. `frontend` - administrative dashboard to view the processed data by administrator.

## Technological stack

Agent: Golang

Storage service: Golang, REST API

Processing service: Python, REST API

Database: PostgreSQL (or SQLite for tests)

Frontend service: Vue.js


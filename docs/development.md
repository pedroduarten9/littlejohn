# Overview

The aim of this file is to document development guidelines.

## Go Folder structure

 - /cmd - this folder is where the executables lie, for this project only `api.go` exists
 - /internal - this folder is where the internal logic lies.
   - /api - This folder is where api handling related logic lies, from endpoints definitions to middlewares...
   - /domain - This is where domain logic lies, generation and static information
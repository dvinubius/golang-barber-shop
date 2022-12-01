# Barber Shop in Go

The classic CS problem by Dijkstra.

## Requirements 

Customers come to the barber shop during business hours.
If waiting room (fixed capacity) is not full, a customer enters the waiting room
If waiting room is full, customer goes somewhere else.
The barber is asleep as long as there are no customers

## Simulation

First the shop opens, then customers begin to come.
The shop has a set time for business hours. After that time passes, no more customers attempt to visit.

Solved via **channels** and a **mutex**.

Pretty terminal output 🤩

![Screenshot 2022-12-01 at 23 36 25](https://user-images.githubusercontent.com/32189942/205164878-1a4ed80d-1fdd-493a-8600-7cb4ea154304.png)

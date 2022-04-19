# Food Order API

## Summary

This is an api that stores and updates food orders that can be queried

## Build

This backend uses _Fiber_ to setup the api routes and _Gorm_ to manage the sqlite3 database interactions

## Routes

- Get Routes: /orders/list -> /orders/complete -> /orders/incomplete
- Post Routes: /orders/add -> /orders/completion -> /orders/update -> /orders/delete -> /orders/current

## Final Comments and Notes

- The admin panel at the base route has no interactivity implemented
- The database file is called _orders.db_ and has a single table called _food_orders_

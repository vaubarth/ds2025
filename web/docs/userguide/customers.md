---
sidebar_position: 3
---
# Customers

A customer is the base entity for all our services. A customer can have any number of services.

All services require a customer, for internal services the customer is `DS`.

## Listing customers

To list all customers with `ds` run
```bash
ds customers

> Listing all known customers
+----+------------+
| ID | CUSTOMER   |
+----+------------+
|  1 | Customer A |
|  1 | Customer B |
|  1 | Customer C |
|  1 | Customer D |
|  1 | Customer E |
+----+------------+
```

## Adding a new customer

Creating a new customer will set up everything that is needed to later add services to that customer.

It will create all required namespaces and add the customer to the list of known customers.
```bash
ds customers new CustomnerName

> Adding new customer CustomnerName
Creating customer data ... done! [100 in 1.003s]
Registering namespaces ... done! [100 in 1.003s]
```
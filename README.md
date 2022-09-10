# Dining Hall

## Description

The ```Dining hall``` has a finite amount of ```tables``` that "clients" can occupy. For simplicity's sake, at any given time a ```table```
can have only one order, thus if a restaurant has 6 ```tables``` occupied from total number of 10 tables , it can at most have 6
orders.

Main work unit of the ```Dining hall``` are ```waiters``` (which are a bit counter intuitively named) that take orders from the
```tables``` (clients).
The purpose of ```waiters``` is to "find" ```orders``` . Restaurant has limited number of ```waiters``` and is less than number of ```tables``` .
For example at restaurant with 10 tables we can have 4-5 waiters.

## Restaurant Menu

```Tables``` (clients) generates ```orders``` based on restaurant menu. Menu consist of ```foods``` .
```Foods``` used in our simulation are given in the list bellow:

- pizza, id=1

``` json
{
	"id": 1,
	"name": "pizza",
	"preparation-time": 20 ,
	"complexity": 2 ,
	"cooking-apparatus": "oven"
}
```
- salad, id=2

``` json
{
	"id": 2,
	"name": "salad",
	"preparation-time": 10 ,
	"complexity": 1 ,
	"cooking-apparatus": null
}
``` 

- zeama, id=3

``` json
{
	"id": 3,
	"name": "zeama",
	"preparation-time": 7 ,
	"complexity": 1 ,
	"cooking-apparatus": "stove"
}
``` 
(See full list here)


## Order

```Tables``` (clients) generates ```orders``` . An order should contain the following information:

- unique order id;
- one or more menu items where the items indicate the ids of the menu items;
- the priority of the order (where it ranges from 1 to 5 , 1 being the smallest priority, and 5 - with the highest one);
- maximum wait time that a client is willing to wait for its order and it should be calculated by taking the item with the highest preparation-time from the order and multiply it by 1.3.

An example of an order:
``` json
{
	"id": 1,
	"items": [ 3, 4, 4, 2 ],
	"priority": 3 ,
	"max_wait": 45
}
```
**The timer of an order starts from the moment it's created.** 

```Order``` has to be picked up by a waiter. The time it takes for a waiter varies, and I would say that a time between 2 and 4
should be realistic enough.

## Additional Info

The ```Dinning Hall``` consists of ```tables``` and ```waiters``` . We have to design a mechanism which will simulate tables occupation.
At start of simulation, tables should not be totally occupied and we have to take into a count that it takes time for a table to be occupied after it was vacated.

In the ```Dinning Hall``` we should have a collection(array) of ```tables``` .
```Tables``` should be dedicated objects. Each ```table``` should have a state of:

- being free
- waiting to make a order
- waiting for a order to be served

```Waiters``` should be object instances which run their logic of serving tables on separate ```threads``` , one thread per ```waiter``` .
```Waiters``` should look for tables which were not served, meaning that order was not picked up yet. For ```Waiters``` which are running on separate ```threads``` , tables are shared resource. ```Waiters``` are looking in the collection of ```tables``` for such table which is ready to make a order. When ```waiter``` is picking up the order from a ```table``` , it(table) should generate a random order with random foods and random number of foods, random priority and unique order ID.

Number of ```tables``` and ```waiters``` should be configurable.

After picking up an ```order``` , don't forget that this operation takes some amount of time. ```Waiter``` have to send ```order``` to
kitchen by performing HTTP (POST) request, with order details.

When ```order``` will be ready, ```kitchen``` will send a HTTP (POST) request back to ```Dinning Hall``` . Our ```Dinning Hall``` server has
to handle that request and to notify ```waiter``` that order is ready to be served to the table which requested this order.
Our task here is to design a mechanism for serving prepared ```orders``` to ```tables``` . The order should be served to the ```table```
by the ```waiter``` which picked up that specific order. When ```order``` is served ```table``` should check that served ```order``` is the
same ```order``` which was requested.

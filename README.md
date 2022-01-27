# myprogram

- How to run programe with given poolsize and number of jobs

    jobs=5 pool=3 go run main.go 
   
- O/P
------
created queue
worker 1  job 1 produced value-----> {"worker":1,"job":1}
worker 3 consuming  data------> {"worker":1,"job":1}
worker 2  job 2 produced value-----> {"worker":2,"job":2}
worker 2 consuming  data------> {"worker":2,"job":2}
worker 3  job 3 produced value-----> {"worker":3,"job":3}
worker 1 consuming  data------> {"worker":3,"job":3}
worker 1  job 4 produced value-----> {"worker":1,"job":4}
worker 3 consuming  data------> {"worker":1,"job":4}
worker 2  job 5 produced value-----> {"worker":2,"job":5}
worker 2 consuming  data------> {"worker":2,"job":5}

------
   

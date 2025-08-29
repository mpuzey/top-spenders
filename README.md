# Top Spenders 

## Running the top spenders 

Run the top spenders aggregation tool for your desired month and year combination:

```
cd cmd/top-spenders 
// For results from Febuary 2020
go run . -month 2 -year 202
```

The output will look similar to the following:
```
go run .                                                                                                                   ✔  at 16:21:27 
Rank 1: Yahya Frey (yahya.frey@mailinator.com) - £93.28 total spent
Rank 2: Taliah Murillo (taliah.murillo@mailinator.com) - £93.10 total spent
Rank 3: Amiya Mays (amiya.mays@mailinator.com) - £91.19 total spent
Rank 4: Vivek Talley (vivek.talley@mailinator.com) - £87.01 total spent
Rank 5: Bronte Parkes (bronte.parkes@mailinator.com) - £85.91 total spent
```

## Reading the docs

First install godoc
```
go install golang.org/x/tools/cmd/godoc@latest
```

Run godoc across the project

```
godoc -http=:6060
```
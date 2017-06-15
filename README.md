Database spikes for customise my data
===================================

The input files used in the tests are zipped in the input-files directory.

### Test queries 

Queries are in 'pseudo' sql due to variances in the databases. They are provided only to show the variations on filters

#### ASHE07E dataset

##### select all data

```
SELECT * from observation
```

##### select a single point (filter on all dimensions)

```
SELECT * from observation
WHERE Geography="K02000001" 
AND Year="2015" 
AND Sex="CI_0006618" 
AND `Working pattern`="CI_0006618" 
AND Earnings="CI_0021537" 
AND `Earnings statistics`="CI_0006603"
```

##### select a single dimension value

```
SELECT * from observation
WHERE `Earnings statistics`="CI_0006603"
```

##### select multiple dimension values

```
SELECT * from observation
WHERE `Earnings statistics`="CI_0006603" 
OR `Earnings statistics`="CI_0006604"
```

##### select first 3 rows returned from single dimension values

```
SELECT * from observation
WHERE `Earnings statistics`="CI_0006603"
LIMIT 3
```

##### select rows 5-8 returned from single dimension values

```
SELECT * from observation
WHERE `Earnings statistics`="CI_0006603"
LIMIT 4
OFFSET 4
```

##### select first 3 rows after sorting on returned data from single dimension values

```
SELECT * from observation
WHERE `Earnings statistics`="CI_0006603"
ORDER BY Dimension_Value_4 DESC
LIMIT 3
```

##### select rows 5-8 after sorting on returned data from single dimension values

```
SELECT * from observation
WHERE `Earnings statistics`="CI_0006603"
ORDER BY Dimension_Value_4 DESC
LIMIT 4
OFFSET 4
```

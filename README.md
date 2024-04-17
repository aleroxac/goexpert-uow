# goexpert-uow


## Modo de uso
``` shell
make compose_up
make generate
make test
```

## UOW - Unit of work
``` 
### Example of a unit of work
{
    BEGIN
    Transaction_1 -> Repository_1
    Transaction_2 -> Repository_2
    COMMIT / ROLLBACK
}

### flow
Repositories > UOW > getRespository > Respository(Transaction)

### Unit of work functions
Register
Get
Do
CommitOrRollback
Rollback
UnRegister
```

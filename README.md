trmutex
=======

Trmutex means "temporary resource mutex" and therefore is designed to manipulate synchronization of temporary resources which might not exist before. 

The difference between trmutex.Mutex and sync.Mutex is that the former is able to create an underlying mutex if it does not exist before and to synchronize the following access launched by other goroutines, and to destroy the underlying mutex when all of the resource access synchronization is finished.

Trmutex is composed of two structs, Mutex, which is to synchronize like sync.Mutex, and Factory, used to create a mutex corresponding to a temporary resource specified with a string id.

Trmutex is simple and easy to use. Here is an example:


    factory := trmutex.NewFactory()
    mutex := factory.Require("MyTemporaryResourceName")
    mutex.Lock()
    defer mutex.Lock()
    ...


In addition, trmutex.Mutex implements sync.Locker and is free to be copied without causing any deallock.

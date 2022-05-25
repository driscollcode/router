# Mockgen Testing

You can use mockgen to auto generate a mock request which can make testing handlers very easy. 

## Generating Mocks

To generate mocks, create an interface which wrappers the ``router.Request`` interface and gives a
``go generate`` instruction. Have a look at this example:

```go
//go:generate mockgen -destination=mock-request.go -package=mock . Request
type Request interface {
	router.Request
}
```

When you run ``go generate`` a new mock request will be created in the file ``mock-request.go``. The
mock allows you to set expectations for which methods will be called and you can define the return 
values.

You can see a working example of this in the ``mock`` folder. To make it easier for you to generate
your mock request, you can use the command ``make mocks``.

## Testing With Mocks

Mock requests make it very easy to test handlers - have a look at this example:

```go
When("a required argument - userId - is not supplied", func() {
    It("should return an error", func() {
        mockRequest.EXPECT().ArgExists("userId").Return(false)
        mockRequest.EXPECT().Error("No user Id supplied")

        myHandler.Handler(mockRequest)
    })
})
```

In this example we are instructing our mock request that our handler should call the ``ArgExists`` 
method with the parameter ``"userId"`` - and when it does, the method should return ``false``. 
We are also instructing our mock request that the handler should call the ``Error`` method with the 
argument ``"No user Id supplied"``.

If our handler fails to make either of these calls with the expected parameters, the test fails.

## Full Example

To look at a working example of testing with and generating mocks, look at the file ``unit_test.go``
in this folder. This is a working test example which unit tests a sample handler located in the 
``my-handler`` folder. You can run the tests with the command ``make test``.

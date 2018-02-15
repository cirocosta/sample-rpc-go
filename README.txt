

The net/rpc package is frozen and is not accepting new features.


        >       Package rpc provides access to the exported methods
                of an object across a network or other I/O connection.


        >       A server registers an object, making it visible as a service
                with the name of the type of the object.

        >       After registration, exported methods of the object will be
                accessible remotely.


        >       A server may register multiple objects (services) of different
                types but it is an error to register multiple objects of the same
                type.

                It must satisfy the contraints:

                        - the method's type is exported.
                        - the method is exported.
                        - the method has two arguments, both exported (or builtin) types.
                        - the method's second argument is a pointer.
                        - the method has return type error.

                meaning that it must have the following signature:


                        func (t *T) MethodName(argType T1, replyType *T2) error


                where T1 and T2 can be marshaled by `encoding/gob`:


                        Package gob manages streams of gobs - binary values exchanged
                        between an Encoder (transmitter) and a Decoder (receiver).

                        A typical use is transporting arguments and results of remote
                        procedure calls (RPCs) such as those provided by package "net/rpc".

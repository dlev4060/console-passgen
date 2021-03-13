# Console Password generator

You can generate a random password with special characters and 
the required length, or you can use the standard generation parameters.

# Args

--l or --length // an explicit indication of the length of the password. 
Default 12 characters

--spec-disable // Disables the use of special characters when generating

# Example

 // 1
$ ./test --l 25 
Generated password: EhrRrZMmQ8LUmIkiXouBn&stc

 // 2
$ ./test --l 25 --spec-disable
Generated password: Fc2O23DSeGVHsDGWbnkzBV4W8

 // 3
$ ./test
Generated password: E2nrG5VdQ6Wi
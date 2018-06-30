#/bin/bash


go build topN.go
#testcase 1
fail=0
./topN -input ./testcase1.txt -N 10 > result
if diff result out_case1.txt > /dev/null ; then
	echo "Testcase 1 passed."
else
	fail=$((fail+1))
	echo "Testcase 1 failed."
fi
rm result

#testcase 2 
./topN -input ./testcase2.txt -N 10 > result
if diff result out_case2.txt > /dev/null ; then
	echo "Testcase 2 passed."
else
	fail=$((fail+1))
	echo "Testcase 2 failed."
fi
rm result
if [ $fail -eq 0 ]; then
	echo "All tests passed successfully!"
else
	echo "$fail tests failed!"
fi



Task 2.

1. What happens if you switch the order of the statements wgp.Wait() and close(ch) in the end of the main function?

Answer: By closing the channel first, you do not allow all the producers to send to the channel before closing. Therefore, the program will try to send to a closed channel, causing a run-time panic as sending to closed channel is an error.

Receiving from a closed channel is not a run-time panic, that is why we can be sure that the Produce function will be causing the problem with the closed channel.

2. What happens if you move the close(ch) from the main function and instead close the channel in the end of the function Produce?

Answer: Because we aren't sure that all producers have finished sending to the channel when we close it, the program results in a run-time panic as a remaining Produce thread tries to send to a closed channel. All of the Produce threads run concurrently, but when the first one that we created is finished and closed the channel, all the remaining ones will be sending to a closed channel causing a run-time panic.

3. What happens if you remove the statement close(ch) completely?

Answer: The program will function as normal because of the built-in garbage collector Golang has. The garbage collector will take care of the unclosed channel. Close is only an indicator that the channel is closed, and thus not a must.

4. What happens if you increase the number of consumers from 2 to 4?

Answer: The program will run faster. This is because we can have more consumers that can read the channel and output the data that the producers have made. We can expect this behavior as we will have 4 concurrent threads reading from the channel, instead of only 2. Thus speeding up the processes of reading the entire channel.

5. Can you be sure that all strings are printed before the program stops?



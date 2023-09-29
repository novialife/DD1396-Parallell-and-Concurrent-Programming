1. What happens if you remove the go-command from the Seek call in the main function?

Answer: By removing the go-command, seperate go routines are not created, and thus each function call is run after the previous one is finished, instead of concurrently.

2. What happens if you switch the declaration wg := new(sync.WaitGroup) to var wg sync.WaitGroup and the parameter wg *sync.WaitGroup to wg sync.WaitGroup?

Deadlock! When passing a parameter variable with *, you are referencing to the pointer to that variable. When passing the variable directly without *, you are creating a copy of the variable. Therefore, by removing the *, the WaitGroups operate on copies of themselves instead of the actual WaitGroup. That is why we end up in a deadlock.

3. What happens if you remove the buffer on the channel match?

Deadlock! By buffering the channel, we do not require the same amount of send and receives from the channel. The buffer allows us to get around this and avoiding a potential deadlock.

4. What happens if you remove the default-case from the case-statement in the main function?

A deadlock may arise. If the channel match has no send operations, the print can not execute without causing a deadlock. To avoid this, the default case can be implemented to catch when the scenario of no pending send operations is occurring to prevent deadlock. However, in this program nothing will happen if you remove it as the default case will never occur.
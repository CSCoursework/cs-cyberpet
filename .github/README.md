# cyberpet

### No, this isn't really complete.

The current state of this repo pretty much satisfies the project brief with the exception of the minigame, which is incomplete. It does some funky display stuff, but you can't actually play a game in it.

### Yes, there are better ways to do some things.

One of the big things I'm not happy with in this is the extensive use of global variables in the `ui` package. I'd switch it all round to being based on a single struct with bound methods if I had time, but I do not.

This also would have benefitted with more extensive planning and research into how I was going to do certain things. The UI code in this (which is by far the majority of the code in this project) is a little bit all over the place.

### Import graph

![import graph](https://raw.githubusercontent.com/CSCoursework/cs-cyberpet/master/.github/importgraph.png)

### The stat update loops

Statistics are constantly updated in the background while the program is running through the use of goroutines. For this, there are two main functions - `ui.StartStatLoop` and the goroutine that's started by `pet.NewPet`.

The goroutine in `pet.NewPet` serves only to update stats on a fixed interval. It's bound to a specific instance of `pet.Pet`, and on every set interval (currently that's every 5 seconds) it performs the following actions: 

*  applies a delta to each statistic value (so they increase/decrease at a constant rate, with the exception of health)
* checks if any values no longer satisfy the inequality `0<=n<=100`, in which case they are set to either `0` or `100`.
* checks to see if certain statistic values are at maximum
  * If any are, a delta is applied to the health value depending on which statistic value is at maximum
* sends a "signal" through a channel to the UI layer of the app

This signal does two things - it informs the UI update loop that a new change has been made to the stats and they should be updated on screen, and it also informs the UI loop if the pet has died or not. This is done using a boolean.

This boolean value is received in `display.StartStatLoop` is used to control a `for` loop. While `true` values are being sent from the stat update goroutine (which indicates the pet is still alive), the UI will update whenever a value is received. The first time a `false` is received, the `for` loop stops and that goroutine moves onto declaring to the user that the pet has died and then shutting the program down.
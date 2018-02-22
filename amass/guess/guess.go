// Copyright 2017 Jeff Foley. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package guess

type Guesser interface {
	Train(good, bad []string)
	NextGuess() (string, error)
}

// Code generated by "go generate gonum.org/v1/gonum/unit”; DO NOT EDIT.

// Copyright ©2014 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unit

import (
	"errors"
	"fmt"
	"math"
	"unicode/utf8"
)

// Charge represents an electric charge in Coulombs.
type Charge float64

const (
	Yottacoulomb Charge = 1e24
	Zettacoulomb Charge = 1e21
	Exacoulomb   Charge = 1e18
	Petacoulomb  Charge = 1e15
	Teracoulomb  Charge = 1e12
	Gigacoulomb  Charge = 1e9
	Megacoulomb  Charge = 1e6
	Kilocoulomb  Charge = 1e3
	Hectocoulomb Charge = 1e2
	Decacoulomb  Charge = 1e1
	coulomb      Charge = 1.0
	Decicoulomb  Charge = 1e-1
	Centicoulomb Charge = 1e-2
	Millicoulomb Charge = 1e-3
	Microcoulomb Charge = 1e-6
	Nanocoulomb  Charge = 1e-9
	Picocoulomb  Charge = 1e-12
	Femtocoulomb Charge = 1e-15
	Attocoulomb  Charge = 1e-18
	Zeptocoulomb Charge = 1e-21
	Yoctocoulomb Charge = 1e-24
)

// Unit converts the Charge to a *Unit
func (ch Charge) Unit() *Unit {
	return New(float64(ch), Dimensions{
		CurrentDim: 1,
		TimeDim:    1,
	})
}

// Charge allows Charge to implement a Charger interface
func (ch Charge) Charge() Charge {
	return ch
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension
func (ch *Charge) From(u Uniter) error {
	if !DimensionsMatch(u, coulomb) {
		*ch = Charge(math.NaN())
		return errors.New("Dimension mismatch")
	}
	*ch = Charge(u.Unit().Value())
	return nil
}

func (ch Charge) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", ch, float64(ch))
			return
		}
		fallthrough
	case 'e', 'E', 'f', 'F', 'g', 'G':
		p, pOk := fs.Precision()
		w, wOk := fs.Width()
		const unit = " C"
		switch {
		case pOk && wOk:
			fmt.Fprintf(fs, "%*.*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), p, float64(ch))
		case pOk:
			fmt.Fprintf(fs, "%.*"+string(c), p, float64(ch))
		case wOk:
			fmt.Fprintf(fs, "%*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), float64(ch))
		default:
			fmt.Fprintf(fs, "%"+string(c), float64(ch))
		}
		fmt.Fprint(fs, unit)
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%g C)", c, ch, float64(ch))
	}
}

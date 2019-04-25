# fukushima_lambertw
[![Build Status](https://travis-ci.org/jtejido/fukushima_lambertw.svg?branch=master)](https://travis-ci.org/jtejido/fukushima_lambertw) 
[![codecov](https://codecov.io/gh/jtejido/fukushima_lambertw/branch/master/graph/badge.svg)](https://codecov.io/gh/jtejido/fukushima_lambertw)

Toshio Fukushima's method for the Lambert W function.

This code is based on the following publication and its author's fortran code:

***Toshio Fukushima***, **"Precise and fast computation of Lambert W-functions without
transcendental function evaluations"**, *J. Comp. Appl. Math. 244 (2013) 77-89.*

## Usage:

result := fukushima_lambertw.LambertW(branch int, x float64)

**branch** = 0 or -1

Alternatively:

**Wp(W0)** and **Wm(W-1)** is exposed via:

w0 := fukushima_lambertw.LambertW0(x float64)

wm := fukushima_lambertw.LambertWm1(x float64)

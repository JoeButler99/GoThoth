package Thoth

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type ConstantPool struct {
	InitialMin, InitialMax float64
	ConstantSet            []float64
}

type FitnessCases struct {
	TotalCases, Terminals, NumVars, NumConsts                       uint
	ConstLower, ConstUpper, TargetScore, ScalingLower, ScalingUpper float64
	ScalingEnabled                                                  bool
	Multipliers, Targets, CliCase                                   []float64
	Cases                                                           [][]float64
	InputFile                                                       string
	ConstantPool                                                    ConstantPool
}

func (f *FitnessCases) LoadFile(inputFile string) {
	file, err := os.Open(inputFile)
	check_err(err)
	defer file.Close()

	line_no := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bits := strings.Split(scanner.Text(), " ")

		if line_no == 0 {
			// This is the first line of the fitness case with: NUM_VARS >> NUM_CONSTS >>  CONST_LOWER >> CONST_UPPER
			if len(bits) != 4 {
				panic("Bad Fitness case line 0")
			}

			tUint, err := strconv.ParseUint(bits[0], 10, 32)
			check_err(err)
			f.NumVars = uint(tUint)
			tUint, err = strconv.ParseUint(bits[1], 10, 32)
			check_err(err)
			f.NumConsts = uint(tUint)
			f.ConstLower, err = strconv.ParseFloat(bits[2], 64)
			check_err(err)
			f.ConstUpper, err = strconv.ParseFloat(bits[3], 64)
			check_err(err)
			f.Terminals = f.NumVars + f.NumConsts

		} else if line_no == 1 {
			// This is the second line of the fitness cases with: TARGET_SCORE >> TOTAL_CASES >> SCALING_LOWER >> SCALING_UPPER
			if len(bits) != 4 {
				panic("Bad Fitness case line 1")
			}
			f.TargetScore, err = strconv.ParseFloat(bits[0], 64)
			check_err(err)
			tUint, err := strconv.ParseUint(bits[1], 10, 64)
			check_err(err)
			f.TotalCases = uint(tUint)
			f.ScalingLower, err = strconv.ParseFloat(bits[2], 64)
			check_err(err)
			f.ScalingUpper, err = strconv.ParseFloat(bits[3], 64)
			check_err(err)
		} else {
			var caseLine []float64
			for x, val := range bits {
				if x < int(f.NumVars) {
					t, err := strconv.ParseFloat(val, 64)
					check_err(err)
					caseLine = append(caseLine, t)
				} else if x == int(f.NumVars) {
					// The target should be the last element in the line
					t, err := strconv.ParseFloat(val, 64)
					check_err(err)
					f.Targets = append(f.Targets, t)
				} else {
					panic("Too many vars to unpack in fitness case")
				}
			}
			if f.ScalingEnabled {
				panic("Scaling not yet implemented")
			} else {
				f.Multipliers = append(f.Multipliers, 1)
			}
			f.Cases = append(f.Cases, caseLine)
		}

		line_no++
	}
	check_err(scanner.Err())

}

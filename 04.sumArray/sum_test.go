package sum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSumArr(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	got := SumArr(numbers)
	want := 15

	if got != want {
		t.Errorf("Expected %d but got %d given %v", want, got, numbers)
	}
}

func TestSumShort(t *testing.T) {
	t.Run("Checking SumShort first check", func(t *testing.T) {
		numbers := []int{1, -5, 3, 4}
		got := SumShort(numbers)
		want := 3
		if got != want {
			t.Errorf("\nGot: %d \nExpected: %d \nGiven: %v", got, want, numbers)
		}
	})
	t.Run("Checking SumShort second check", func(t *testing.T) {
		numbers := []int{1, 0, 3, 4, 7}
		got := SumShort(numbers)
		want := 15
		if got != want {
			t.Errorf("\nGot: %d \nExpected: %d \nGiven: %v", got, want, numbers)
		}
	})
}

func TestSumAllSlices(t *testing.T) {
	got := SumAllSlices([]int{1, 5}, []int{4, 3, -1})
	want := []int{6, 6}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("\nGot: %v \nExpeced: %v", got, want)
	}
}
func TestSumAllSlicesAlt(t *testing.T) {
	got := SumAllSlicesAlt([]int{-1, 4, 5}, []int{4, 3, -1, 10})
	want := []int{8, 16}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("\nGot: %v \nExpeced: %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	nums1 := []int{1, 2, 3, -2}
	nums2 := []int{-5, 3, -1, 3, 7, -2}
	nums3 := []int{-5, 3, -1, 3, 7, -2, 7, -5}
	nums4 := []int{}
	nums5 := []int{1, 3}

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("\nGot: %v \nWant: %v ", got, want)
		}
	}

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails(nums1, nums2, nums3)
		want := []int{3, 10, 12}
		checkSums(t, got, want)
	})

	t.Run("summation of empty slices", func(t *testing.T) {
		got := SumAllTails(nums4, nums5)
		want := []int{0, 3}
		checkSums(t, got, want)
	})

}

func ExampleSumArr() {
	numbers := [5]int{1, 2, 3, 0, 0}
	sum := SumArr(numbers)
	fmt.Println(sum)
	// Output: 6
}

func ExampleSumShort() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := SumShort(numbers)
	fmt.Println(sum)
	// Output: 15
}

func ExampleSumAllSlices() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{-5, 3, -1, 3, 7}
	res := SumAllSlices(nums1, nums2)
	fmt.Println(res)
	// Output: [6 7]
}

func ExampleSumAllSlicesAlt() {
	nums1 := []int{1, 2, 3, -2}
	nums2 := []int{-5, 3, -1, 3, 7, -2}
	res := SumAllSlicesAlt(nums1, nums2)
	fmt.Println(res)
	// Output: [4 5]
}

func ExampleSumAllTails() {
	nums1 := []int{1, 2, 3, -2}
	nums2 := []int{-5, 3, -1, 3, 7, -2}
	nums3 := []int{-5, 3, -1, 3, 7, -2, 7, -5}
	res := SumAllTails(nums1, nums2, nums3)

	nums4 := []int{1}
	nums5 := []int{5, 6, 4}
	nums6 := []int{}
	res2 := SumAllTails(nums4, nums5, nums6)
	fmt.Println(res, res2)
	// Output: [3 10 12] [0 10 0]
}

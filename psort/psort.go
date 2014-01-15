package psort

import (
	"math"
	"sort"
)

type Interface interface {
	CreateCopy(length int) Interface
	ReferenceFrom(src Interface, p,open_r int) 
	SetValue(a Interface,i int, b Interface, j int)
	GetValue(i int) interface{}
	Len() int
	Swap(i,j int) 
	Less(i,j int) bool
	Less2(a,b interface{}) bool
}

func Psort(a Interface) {
	b := a.CreateCopy(a.Len())
	pmergesort(a,0,a.Len()-1,b,0, nil)
	a.ReferenceFrom(b, 0, a.Len())
}

func insertionSort(a Interface, p,r int) {
	for i:=p+1;i<=r;i++ {
		for j:=i;j>p && a.Less(j,j-1);j-- {
				a.Swap(j,j-1)
		}
	}
}

func pmergesort(a Interface, p,r int, b Interface, s int, isDone chan bool) {
	n := r-p+1
	if n==1 {
		//b[s]=a[p]
		b.SetValue(b,s,a,p)
	} else if n<=1000 {
		t := a.CreateCopy(n)
		t.ReferenceFrom(a,p,r+1)
		sort.Sort(t)
		for i:=0;i<n;i++ {
			b.SetValue(b,s+i,t,i)
		}
	}else {
		t := a.CreateCopy(n)
		q1 := int(math.Ceil(float64((p+r)/2)))
		q2 := q1-p
		isChildDone1 := make(chan bool)
		go pmergesort(a,p,q1,t,0, isChildDone1)
		pmergesort(a,q1+1,r,t,q2+1, nil)
		<-isChildDone1
		pmerge(t,0,q2,q2+1,n-1,b,s, nil)
	}
	if isDone!=nil {
		isDone <- true
	}
}

func pmerge(a Interface, p1,r1,p2,r2 int, b Interface, s int, isDone chan bool) {
	n1 := r1-p1+1
	n2 := r2-p2+1
	if n1<n2 {
		p1,p2 = p2,p1
		r1,r2 = r2,r1
		n1,n2 = n2,n1
	}
	if n1==0 {
		if isDone!=nil {
			isDone <- true
		}
		return
	} else {
		q1 := int(math.Ceil(float64(p1+r1)/2))
		//t := a[p2:r2+1]
		t := a.CreateCopy(r2-p2+1)
		t.ReferenceFrom(a,p2,r2+1)
		count := sort.Search(t.Len(), func(i int) bool {
			return t.Less2(a.GetValue(q1),t.GetValue(i))})
		q2 := p2+count
		q3 := s+q1-p1+q2-p2
		//b[q3]=a[q1]
		b.SetValue(b,q3,a,q1)
		isChildDone1 := make(chan bool)
		go pmerge(a,p1,q1-1,p2,q2-1,b,s, isChildDone1)
		pmerge(a,q1+1,r1,q2,r2,b,q3+1, nil)
		<-isChildDone1
	}
	if isDone!=nil {
		isDone <- true
	}
}

type IntSlice []int
func (owner IntSlice) CreateCopy(length int) Interface {
	return make(IntSlice, length)
}
func (owner IntSlice) ReferenceFrom(src Interface, p,open_r int) {
	//return src.(IntSlice)[p:open_r]
	copy(owner,src.(IntSlice)[p:open_r])
}
func (owner IntSlice) SetValue(a Interface,i int, b Interface,j int) {
	a.(IntSlice)[i]=b.(IntSlice)[j]
}
func (owner IntSlice) GetValue(i int) interface{} {
	return owner[i]
}
func (owner IntSlice) Len() int {
	return len(owner)
}
func (owner IntSlice) Swap(i,j int) {
	owner[i],owner[j] = owner[j],owner[i]
}
func (a IntSlice) Less(i,j int) bool {
	return a[i]<a[j];
}
func (a IntSlice) Less2(x,y interface{}) bool {
	return x.(int)<y.(int)
}

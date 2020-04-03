# timezone
时区
==

<pre>
t, _ := timezone.Parse("2006-01-02 15:04:05", "2019-01-01 23:59:59")
s1 := timezone.Format("2006-01-02 15:04:05", t.Unix())
fmt.Println(s1) // Local

timezone.SetUtc()
s2 := timezone.Format("2006-01-02 15:04:05", t.Unix())
fmt.Println(s2) // UTC

_ = timezone.SetNewYork()
s3 := timezone.Format("2006-01-02 15:04:05", t.Unix())
fmt.Println(s3) // NewYork

_ = timezone.SetShanghai()
s4 := timezone.Format("2006-01-02 15:04:05", t.Unix())
fmt.Println(s4) // Shanghai
</pre>
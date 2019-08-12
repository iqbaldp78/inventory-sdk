package jose

import (
	"testing"
)

func init() {
	Setup("8BF5B017E1C6560647DA276C1BF6391A4FF958911D1056B14C0D1165689985CF", "api.boilerplate.com")
}

func TestExtract(t *testing.T) {
	input := []string{
		`eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjg2MzUxNjg3NzAsImlhdCI6MTU2MTUyNzgzNSwiaXNzIjoiYXBpLmJvaWxlcnBsYXRlLmNvbSIsInBheSI6IntcInBheWxvYWRcIjpcImV5SnBaQ0k2TVN3aVozSnZkWEJmYVdRaU9qRXNJbkp2YkdVaU9udDlmUVwiLFwicHJvdGVjdGVkXCI6XCJleUpoYkdjaU9pSklVekkxTmlKOVwiLFwic2lnbmF0dXJlXCI6XCJmVERKbWpCQmdGZU9JcXN3U3dfQ1h0S2NURnZ6UVVIQTRVZjl6M3FNdTVZXCJ9Iiwic3ViIjoiYWNjZXNzIGdyYW50ZWQgZm9yIDYwMDAwMDAwMDAwIGhvdXIocykifQ.t4iO45XwlWVvMNuuBHJM-0Mr_HEqrBcfemT8sSMLz1Q`,
		`Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjg2MzUxNjg3NzAsImlhdCI6MTU2MTUyNzgzNSwiaXNzIjoiYXBpLmJvaWxlcnBsYXRlLmNvbSIsInBheSI6IntcInBheWxvYWRcIjpcImV5SnBaQ0k2TVN3aVozSnZkWEJmYVdRaU9qRXNJbkp2YkdVaU9udDlmUVwiLFwicHJvdGVjdGVkXCI6XCJleUpoYkdjaU9pSklVekkxTmlKOVwiLFwic2lnbmF0dXJlXCI6XCJmVERKbWpCQmdGZU9JcXN3U3dfQ1h0S2NURnZ6UVVIQTRVZjl6M3FNdTVZXCJ9Iiwic3ViIjoiYWNjZXNzIGdyYW50ZWQgZm9yIDYwMDAwMDAwMDAwIGhvdXIocykifQ.t4iO45XwlWVvMNuuBHJM-0Mr_HEqrBcfemT8sSMLz1Q`,
		`Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjg2MzUxNjg3NzAsImlhdCI6MTU2MTUyNzgzNSwiaXNzIjoiYXBpLmJvaWxlcnBsYXRlLmNvbSIsInBheSI6IntcInBheWxvYWRcIjpcImV5SnBaQ0k2TVN3aVozSnZkWEJmYVdRaU9qRXNJbkp2YkdVaU9udDlmUVwiLFwicHJvdGVjdGVkXCI6XCJleUpoYkdjaU9pSklVekkxTmlKOVwiLFwic2lnbmF0dXJlXCI6XCJmVERKbWpCQmdGZU9JcXN3U3dfQ1h0S2NURnZ6UVVIQTRVZjl6M3FNdTVZXCJ9Iiwic3ViIjoiYWNjZXNzIGdyYW50ZWQgZm9yIDYwMDAwMDAwMDAwIGhvdXIocykifQ`,
	}
	output := []bool{
		true,
		false,
		true,
	}

	for index := range input {
		result, _ := Extract(input[index])
		if result.Payload.IsInitial() != output[index] {
			t.Errorf("Expected result to be `%v`. Got `%v`", output[index], result.Payload.IsInitial())
		}
	}
}

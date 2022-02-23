package crawler

import (
	mock_crawler "attack-launchpad/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestShouldGetPdf(t *testing.T) {
	number := 10
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClient := mock_crawler.NewMockClient(ctrl)
	pdfFetcher := PdfFetcher{
		Client: mockClient,
	}

	mockClient.EXPECT().Get("https://localhost:9090/internet-bills/10").Times(1).Return([]byte("This is my expected pdf"), nil)

	pdf, err := pdfFetcher.Fetch(number)

	assert.Nil(t, err)

	assert.Equal(t, pdf, []byte("This is my expected pdf"))
}

func TestShouldReturnErrorDuringGetPDF(t *testing.T) {
	number := 10
	pdfFetcher := PdfFetcher{}
	_, err := pdfFetcher.Fetch(number)
	assert.NotNil(t, err)

}

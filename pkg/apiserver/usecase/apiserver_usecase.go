package usecase

import (
	"log"
	"strings"

	_apiRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/api/repository"
	_methodRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/method/repository"
	"github.com/Hajime3778/api-creator-backend/pkg/domain"
)

// APIServerUsecase Interface
type APIServerUsecase interface {
	RequestDocumentServer(httpMethod string, url string) (domain.API, domain.Method)
}

type apiServerUsecase struct {
	apiRepo    _apiRepository.APIRepository
	methodRepo _methodRepository.MethodRepository
}

// NewAPIServerUsecase APIServerUsecaseインターフェイスを表すオブジェクトを作成します
func NewAPIServerUsecase(apiRepo _apiRepository.APIRepository, methodRepo _methodRepository.MethodRepository) APIServerUsecase {
	return &apiServerUsecase{
		apiRepo:    apiRepo,
		methodRepo: methodRepo,
	}
}

// RequestDocumentServer リクエスト情報からMethodを特定し、ドキュメントに対してCRUDします
func (u *apiServerUsecase) RequestDocumentServer(httpMethod string, url string) (domain.API, domain.Method) {
	api, err := u.apiRepo.GetByURL(url)
	if err != nil {
		log.Fatal(err)
	}

	methods, err := u.methodRepo.GetListByAPIID(api.ID)
	if err != nil {
		log.Fatal(err)
	}

	// MethodのURL部分を抽出
	requestedMethodURL := strings.Replace(url, api.URL, "", 1)

	// 区切り文字で分割する
	//params := regexp.MustCompile("[/?]").Split(methodURL, -1)
	requestedSlashCount := strings.Count(requestedMethodURL, "/")

	var returnMethod domain.Method
	// ※パラメータが2つ以上やクエリストリングは現在の仕様にないので今は考えない！
	for _, method := range methods {
		if method.Type == httpMethod {
			// methodにURLがないパターン
			if requestedMethodURL == "" && method.URL == "" {
				returnMethod = method
				break
			}

			// リクエストとMethod.URLの/数が同じものを検索する
			if requestedSlashCount == strings.Count(method.URL, "/") {
				returnMethod = method
				break
			}
		}
	}
	return api, returnMethod
}

package usecase

import (
	"errors"
	"regexp"
	"strings"

	_apiRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/api/repository"
	_methodRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/method/repository"
	"github.com/Hajime3778/api-creator-backend/pkg/domain"
)

// APIServerUsecase Interface
type APIServerUsecase interface {
	RequestDocumentServer(httpMethod string, url string) (domain.Method, string, error)
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
func (u *apiServerUsecase) RequestDocumentServer(httpMethod string, url string) (domain.Method, string, error) {
	api, err := u.apiRepo.GetByURL(url)
	if err != nil {
		return domain.Method{}, "", err
	}

	// 対象のメソッドを取得
	method, err := u.getRequestedMethod(httpMethod, url, api)

	// リクエストされたパラメータを取得
	param := getRequestedURLParameter(url, api.URL, method.URL)

	return method, param, err
}

// getRequestedMethod リクエストされたHTTPメソッド、URL、APIから、対象のMethodを返却します
func (u *apiServerUsecase) getRequestedMethod(httpMethod string, requestedURL string, api domain.API) (domain.Method, error) {

	var returnMethod domain.Method

	methods, err := u.methodRepo.GetListByAPIID(api.ID)
	if err != nil {
		return returnMethod, err
	}

	// MethodのURL部分を抽出
	requestedMethodURL := strings.Replace(requestedURL, api.URL, "", 1)

	// /{}で囲まれた箇所を削除
	re := regexp.MustCompile(`/\{.+?\}`)

	// ※パラメータが2つ以上やクエリストリングは現在の仕様にないので今は考えない！
	for _, method := range methods {
		if method.Type != httpMethod {
			continue
		}

		// methodにURLがないパターン
		if requestedMethodURL == "" && method.URL == "" {
			returnMethod = method
			break
		}

		// メソッド名が指定されているパターン(例："/foo/{bar}")
		// methodのURLからパラメータ部分を削除
		methodURL := re.ReplaceAllString(method.URL, "")
		if methodURL != "" {
			methodURL = methodURL + "/"
		}
		// パラメータ以外のmethodのURL(メソッド名)が合致していれば返却
		if methodURL != "" && strings.Contains(requestedMethodURL, methodURL) {
			returnMethod = method
			break
		} else {
			// パラメータのみ指定されているパターン(例："/{bar}")
			params := re.FindAllStringSubmatch(method.URL, -1)
			paramsCount := len(params)
			methodURLSlushCount := strings.Count(method.URL, "/")
			requestedSlushCount := strings.Count(requestedMethodURL, "/")

			// MethodURL,リクエスト内の"/"の数と、Methodのパラメータの数が同じであれば返却
			if methodURLSlushCount == requestedSlushCount && requestedSlushCount == paramsCount {
				returnMethod = method
				break
			}
		}
	}

	if returnMethod.ID == "" {
		return returnMethod, errors.New("対象のメソッドが見つかりません")
	}

	return returnMethod, err
}

// getRequestedURLParameter リクエストされたURLパラメータを取得します
func getRequestedURLParameter(requestedURL string, apiURL string, methodURL string) string {
	if methodURL != "" {
		// /{}で囲まれた箇所を削除
		re := regexp.MustCompile(`/\{.+?\}`)
		methodURL = re.ReplaceAllString(methodURL, "")
	}

	// リクエストされたパラメータを取得
	param := strings.Replace(requestedURL, apiURL+methodURL, "", 1)

	// 最初の/を削除する
	if param != "" {
		param = param[1:]
	}

	return param
}

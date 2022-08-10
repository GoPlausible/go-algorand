// Package private provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package private

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Aborts a catchpoint catchup.
	// (DELETE /v2/catchup/{catchpoint})
	AbortCatchup(ctx echo.Context, catchpoint string) error
	// Starts a catchpoint catchup.
	// (POST /v2/catchup/{catchpoint})
	StartCatchup(ctx echo.Context, catchpoint string) error
	// Return a list of participation keys
	// (GET /v2/participation)
	GetParticipationKeys(ctx echo.Context) error
	// Add a participation key to the node
	// (POST /v2/participation)
	AddParticipationKey(ctx echo.Context) error
	// Delete a given participation key by ID
	// (DELETE /v2/participation/{participation-id})
	DeleteParticipationKeyByID(ctx echo.Context, participationId string) error
	// Get participation key info given a participation ID
	// (GET /v2/participation/{participation-id})
	GetParticipationKeyByID(ctx echo.Context, participationId string) error
	// Append state proof keys to a participation key
	// (POST /v2/participation/{participation-id})
	AppendKeys(ctx echo.Context, participationId string) error

	// (POST /v2/shutdown)
	ShutdownNode(ctx echo.Context, params ShutdownNodeParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AbortCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) AbortCatchup(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameter("simple", false, "catchpoint", ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AbortCatchup(ctx, catchpoint)
	return err
}

// StartCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) StartCatchup(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameter("simple", false, "catchpoint", ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.StartCatchup(ctx, catchpoint)
	return err
}

// GetParticipationKeys converts echo context to params.
func (w *ServerInterfaceWrapper) GetParticipationKeys(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetParticipationKeys(ctx)
	return err
}

// AddParticipationKey converts echo context to params.
func (w *ServerInterfaceWrapper) AddParticipationKey(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddParticipationKey(ctx)
	return err
}

// DeleteParticipationKeyByID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteParticipationKeyByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameter("simple", false, "participation-id", ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteParticipationKeyByID(ctx, participationId)
	return err
}

// GetParticipationKeyByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetParticipationKeyByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameter("simple", false, "participation-id", ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetParticipationKeyByID(ctx, participationId)
	return err
}

// AppendKeys converts echo context to params.
func (w *ServerInterfaceWrapper) AppendKeys(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameter("simple", false, "participation-id", ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AppendKeys(ctx, participationId)
	return err
}

// ShutdownNode converts echo context to params.
func (w *ServerInterfaceWrapper) ShutdownNode(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty":  true,
		"timeout": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ShutdownNodeParams
	// ------------- Optional query parameter "timeout" -------------
	if paramValue := ctx.QueryParam("timeout"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "timeout", ctx.QueryParams(), &params.Timeout)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timeout: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ShutdownNode(ctx, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.DELETE("/v2/catchup/:catchpoint", wrapper.AbortCatchup, m...)
	router.POST("/v2/catchup/:catchpoint", wrapper.StartCatchup, m...)
	router.GET("/v2/participation", wrapper.GetParticipationKeys, m...)
	router.POST("/v2/participation", wrapper.AddParticipationKey, m...)
	router.DELETE("/v2/participation/:participation-id", wrapper.DeleteParticipationKeyByID, m...)
	router.GET("/v2/participation/:participation-id", wrapper.GetParticipationKeyByID, m...)
	router.POST("/v2/participation/:participation-id", wrapper.AppendKeys, m...)
	router.POST("/v2/shutdown", wrapper.ShutdownNode, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+x9+3PcNtLgv4Ka/aoc+4Yzkl+7VlXqO8VyEl0cR2Up2bvP9iUYsmcGKxJgAFDSxKf/",
	"/QoNgARJcIZ6rHKp80+2hng0Go1Gv/F5koqiFBy4VpODz5OSSlqABol/0TQVFdcJy8xfGahUslIzwScH",
	"/htRWjK+mkwnzPxaUr2eTCecFtC0Mf2nEwm/V0xCNjnQsoLpRKVrKKgZWG9K07oe6SpZicQNcWiHOD6a",
	"XG/5QLNMglJ9KH/i+YYwnuZVBkRLyhVNzSdFLpleE71mirjOhHEiOBCxJHrdakyWDPJMzfwif69AboJV",
	"usmHl3TdgJhIkUMfzteiWDAOHiqogao3hGhBMlhiozXVxMxgYPUNtSAKqEzXZCnkDlAtECG8wKticvBh",
	"ooBnIHG3UmAX+N+lBPgDEk3lCvTk0zS2uKUGmWhWRJZ27LAvQVW5VgTb4hpX7AI4Mb1m5MdKabIAQjl5",
	"/+1r8uzZs1dmIQXVGjJHZIOramYP12S7Tw4mGdXgP/dpjeYrISnPkrr9+29f4/ynboFjW1GlIH5YDs0X",
	"cnw0tADfMUJCjGtY4T60qN/0iByK5ucFLIWEkXtiG9/rpoTz/6m7klKdrkvBuI7sC8GvxH6O8rCg+zYe",
	"VgPQal8aTEkz6Ie95NWnz/vT/b3rv304TP7L/fni2fXI5b+ux92BgWjDtJISeLpJVhIonpY15X18vHf0",
	"oNaiyjOyphe4+bRAVu/6EtPXss4LmleGTlgqxWG+EopQR0YZLGmVa+InJhXPDZsyozlqJ0yRUooLlkE2",
	"Ndz3cs3SNUmpskNgO3LJ8tzQYKUgG6K1+Oq2HKbrECUGrlvhAxf0/y4ymnXtwARcITdI0lwoSLTYcT35",
	"G4fyjIQXSnNXqZtdVuRsDQQnNx/sZYu444am83xDNO5rRqgilPiraUrYkmxERS5xc3J2jv3dagzWCmKQ",
	"hpvTukfN4R1CXw8ZEeQthMiBckSeP3d9lPElW1USFLlcg167O0+CKgVXQMTiX5Bqs+3/4/Snd0RI8iMo",
	"RVdwQtNzAjwV2fAeu0ljN/i/lDAbXqhVSdPz+HWds4JFQP6RXrGiKgivigVIs1/+ftCCSNCV5EMA2RF3",
	"0FlBr/qTnsmKp7i5zbQtQc2QElNlTjczcrwkBb36em/qwFGE5jkpgWeMr4i+4oNCmpl7N3iJFBXPRsgw",
	"2mxYcGuqElK2ZJCRepQtkLhpdsHD+M3gaSSrABw/yCA49Sw7wOFwFaEZc3TNF1LSFQQkMyM/O86FX7U4",
	"B14zOLLY4KdSwgUTlao7DcCIU28Xr7nQkJQSlixCY6cOHYZ72DaOvRZOwEkF15RxyAznRaCFBsuJBmEK",
	"JtyuzPSv6AVV8PL50AXefB25+0vR3fWtOz5qt7FRYo9k5F40X92BjYtNrf4jlL9wbsVWif25t5FsdWau",
	"kiXL8Zr5l9k/j4ZKIRNoIcJfPIqtONWVhIOP/In5iyTkVFOeUZmZXwr7049VrtkpW5mfcvvTW7Fi6Slb",
	"DSCzhjWqTWG3wv5jxouzY30VVRreCnFeleGC0pZWutiQ46OhTbZj3pQwD2tVNtQqzq68pnHTHvqq3sgB",
	"IAdxV1LT8Bw2Egy0NF3iP1dLpCe6lH+Yf8oyN711uYyh1tCxu2/RNuBsBodlmbOUGiS+d5/NV8MEwGoJ",
	"tGkxxwv14HMAYilFCVIzOygtyyQXKc0TpanGkf5DwnJyMPnbvDGuzG13NQ8mf2t6nWInI49aGSehZXmD",
	"MU6MXKO2MAvDoPETsgnL9lAiYtxuoiElZlhwDheU61mjj7T4QX2AP7iZGnxbUcbiu6NfDSKc2IYLUFa8",
	"tQ0fKRKgniBaCaIVpc1VLhb1D18dlmWDQfx+WJYWHygaAkOpC66Y0uoxLp82Jymc5/hoRr4Lx0Y5W/B8",
	"Yy4HK2qYu2Hpbi13i9WGI7eGZsRHiuB2CjkzW+PRYGT4+6A41BnWIjdSz05aMY2/d21DMjO/j+r81yCx",
	"ELfDxIValMOcVWDwl0Bz+apDOX3CcbacGTns9r0d2ZhR4gRzK1rZup923C14rFF4KWlpAXRf7F3KOGpg",
	"tpGF9Y7cdCSji8IcnOGA1hCqW5+1nechCgmSQgeGb3KRnt/DeV+YcfrHDocna6AZSJJRTYNz5c5L/M7G",
	"jt9jP+QIICOC/U/4H5oT89kQvuGLdlijsDOkXxGY1zOj51rp2c5kGqD+LUhhVVtiVNIbQfm6mbzHIyxa",
	"xvCIN1abJtjDL8IsvbGVHS6EvB29dAiBk8YCSKgZNTgu087OYtOqTBx+IlYE26AzUON06QuTIYa6w8dw",
	"1cLCqab/BiwoM+p9YKE90H1jQRQly+EezuuaqnV/EUate/aUnH5/+GL/6a9PX7w0ekkpxUrSgiw2GhT5",
	"yknTROlNDo/7K0N5tsp1fPSXz73dqD1ubBwlKplCQcv+UNYeZS8t24yYdn2stdGMq64BHHMsz8CwF4t2",
	"Yk2tBrQjpsydWCzuZTOGEJY1s2TEQZLBTmK66fKaaTbhEuVGVvehfICUQkYsInjEtEhFnlyAVExEjNsn",
	"rgVxLbxAUnZ/t9CSS6qImRuNdRXPQM5ilKWvOILGNBRq14Vqhz674g1u3IBUSrrpod+uN7I6N++YfWkj",
	"39t+FClBJvqKkwwW1aoluy6lKAglGXbEi+MtW611cI+eSCGW9y5uRGeJLQk/oIGd5KaPu+msbIAAvxMZ",
	"GEWpUvfA3pvBGuwZyglxRhei0oQSLjJArapSccY/4JpDnwC6MnR4l+i1FSwWYCT4lFZmtVVJ0FDfo8Wm",
	"Y0JTS0UJokYNWDJrE7RtZaezbp9cAs2MZA+ciIUzFzpDJi6SopdBe9bprp2IrtOCq5QiBaWMRmbl7J2g",
	"+XaWLPUWPCHgCHA9C1GCLKm8JbBaaJrvABTbxMCt5URnY+1DPW76bRvYnTzcRiqNUmapwAil5sDloGEI",
	"hSNxcgESbY3/1v3zk9x2+6pyIBLAiVZnrEDdjlMuFKSCZyo6WE6VTnYdW9OoJf+ZFQQnJXZSceAB+8Jb",
	"qrS1ODOeoS5g2Q3OYw0PZophgAevQDPyL/7264+dGj7JVaXqq1BVZSmkhiy2Bg5XW+Z6B1f1XGIZjF3f",
	"t1qQSsGukYewFIzvkGVXYhFEdW2YcS6Z/uLQfGHugU0UlS0gGkRsA+TUtwqwG3pDBwAximPdEwmHqQ7l",
	"1C7Y6URpUZbm/Omk4nW/ITSd2taH+uembZ+4qG74eibAzK49TA7yS4tZ6wdfUyO048ikoOfmbkIR3JrG",
	"+zCbw5goxlNItlG+OZanplV4BHYc0gHtx0XaBLN1DkeHfqNEN0gEO3ZhaMEDqtgJlZqlrERJ4gfY3Ltg",
	"1Z0gatIhGWjKjHoQfLBCVhn2J9bX0R3zdoLWKKm5D35PbI4sJ2cKL4w28OewQdvuiXWinwWu93uQFCOj",
	"mtNNOUFAvWvOXMhhE7iiqc435prTa9iQS5BAVLUomNY2KqItSGpRJuEAUYvElhmdTcg6oP0OjDFSneJQ",
	"wfL6WzGdWLFlO3xnHcGlhQ4nMJVC5CNs5z1kRCEYZVsnpTC7zlwQjo/U8JTUAtIJMWgQrJnnI9VCM66A",
	"/C9RkZRyFMAqDfWNICSyWbx+zQzmAqvndFb0BkOQQwFWrsQvT550F/7kidtzpsgSLn3kmmnYRceTJ6gl",
	"nQilW4frHlR0c9yOI7wdTTXmonAyXJenzHbaItzIY3bypDN4bd8xZ0opR7hm+XdmAJ2TeTVm7SGNrKla",
	"7147jjvKChMMHVs37ju6EP89OnwzdAy6/sSB46X5OOR7MfJVvrkHPm0HIhJKCQpPVaiXKPtVLMPgRnfs",
	"1EZpKPqqve3664Bg896LBT0pU/CccUgKwWETjednHH7Ej7He9mQPdEYeO9S3Kza14O+A1Z5nDBXeFb+4",
	"2wEpn9ROx3vY/O64HatOGNaJWinkJaEkzRnqrIIrLatUf+QUpeLgLEdM/V7WH9aTXvsmccUsoje5oT5y",
	"qgwOa1k5ap5cQkQL/hbAq0uqWq1A6Y58sAT4yF0rxknFmca5CrNfid2wEiTa22e2ZUE3ZElzVOv+ACnI",
	"otLtGxOjz5Q2Wpc1MZlpiFh+5FSTHIwG+iPjZ1c4nA/y8jTDQV8KeV5jYRY9DyvgoJhK4i6J7+zX76la",
	"++Wbhp5Jus7WiGLGb0LUNhpa4e3/+6v/PPhwmPwXTf7YS179t/mnz8+vHz/p/fj0+uuv/0/7p2fXXz/+",
	"z/+I7ZSHPRYb5SA/PnLS5PERigyNcakH+4NZHArGkyiRna2BFIxjiG2HtshXRvDxBPS4MVO5Xf/I9RU3",
	"hHRBc5ZRfTty6LK43lm0p6NDNa2N6CiQfq2fYi7dlUhKmp6jR2+yYnpdLWapKOZeip6vRC1RzzMKheD4",
	"LZvTks1VCen8Yn/HlX4HfkUi7KrDZG8tEPT9gfF4RjRZuhBFPHnLiluiqJQzUmK4jvfLiOW0jlm1uWoH",
	"BAMa19Q7Fd2fT1+8nEybQMT6u9HU7ddPkTPBsqtYuGkGVzFJzR01PGKPFCnpRoGO8yGEPeqCsn6LcNgC",
	"jIiv1qx8eJ6jNFvEeeX3jjE6je+KH3MbgGFOIppnN87qI5YPD7eWABmUeh3LYWnJHNiq2U2AjkullOIC",
	"+JSwGcy6Gle2AuWdYTnQJeZSoIlRjAnqqs+BJTRPFQHWw4WMUmti9INisuP719OJEyPUvUv2buAYXN05",
	"a1us/1sL8ui7N2dk7livemQjn+3QQaxqxJLhwrFazjbDzWzmng39/sg/8iNYMs7M94OPPKOazhdUsVTN",
	"KwXyG5pTnsJsJciBj/A6opp+5D2ZbTC5NoitI2W1yFlKzkPZuiFPmzDVH+Hjxw+G43/8+KnnuelLwm6q",
	"KH+xEySXTK9FpROXEZJIuKQyi4Cu6owAHNnmc22bdUrc2JYVu4wTN36c59GyVN3I4P7yyzI3yw/IULm4",
	"V7NlRGkhvVRjRB0LDe7vO+EuBkkvfTpRpUCR3wpafmBcfyLJx2pv7xmQVqjsb054MDS5KaFl87pV5HLX",
	"3oULtxoSXGlJk5KuQEWXr4GWuPsoeRdoXc1zgt1aIbo+oAWHahbg8TG8ARaOG4cb4uJObS+f2htfAn7C",
	"LcQ2RtxonBa33a8gaPfW29UJ/O3tUqXXiTnb0VUpQ+J+Z+qMv5URsrwnSbEVN4fAJUcugKRrSM8hwzwt",
	"KEq9mba6e2elE1k962DK5jPaqEJMukHz4AJIVWbUCfWUb7rZDwq09ikf7+EcNmeiydm5SbpDO/peDR1U",
	"pNRAujTEGh5bN0Z3853jGyOOy9IHsWPApieLg5oufJ/hg2xF3ns4xDGiaEWHDyGCyggiLPEPoOAWCzXj",
	"3Yn0Y8sz+srC3nyR9EfP+4lr0qhhznkdrgaD3u33AjA5WlwqsqBGbhcur9dGmAdcrFJ0BQMScmihHRnH",
	"3bLq4iC77r3oTSeW3Qutd99EQbaNE7PmKKWA+WJIBZWZTsiCn8k6AXAFM4LlOhzCFjmKSXW0hGU6VLYs",
	"5bb+wBBocQIGyRuBw4PRxkgo2ayp8inHmJntz/IoGeDfmDGxLU/uOPC2B+nXdRac57ndc9rTLl22nE+R",
	"83lxoWo5IsfNSPgYABbbDsFRAMogh5VduG3sCaXJ3mg2yMDx03KZMw4kiTnuqVIiZTZnvLlm3Bxg5OMn",
	"hFhjMhk9QoyMA7DRuYUDk3ciPJt8dRMgucs+oX5sdIsFf0M87NKGZhmRR5SGhTM+EFTnOQB10R71/dWJ",
	"OcJhCONTYtjcBc0Nm3MaXzNIL10LxdZOcpZzrz4eEme32PLtxXKjNdmr6DarCWUmD3RcoNsC8XZRIrYF",
	"CvHlbFk1robu0jFTD1zfQ7j6Kkj0uhUAHU2/KYnkNL+dGlr7bu7fZA1LnzYJzD6qNEb7Q/QT3aUB/PVN",
	"EHVq1kn3uo4q6W23azsrLZCfYqzYnJG+r6PvUVGQA0rESUuCSM5jHjAj2AOy21PfLdDcMfeN8s3jwJcv",
	"YcWUhsYWbW4l71x5aNscxZR7IZbDq9OlXJr1vRei5tE2pxM7tpb54Cu4EBqSJZNKJ2jIjy7BNPpWoUb5",
	"rWkaFxTa0QK2+gzL4rwBpz2HTZKxvIrTq5v3hyMz7bvaCKOqxTlsUBwEmq7JAqslRWOItkxtw8y2Lvit",
	"XfBbem/rHXcaTFMzsTTk0p7jL3IuOpx3GzuIEGCMOPq7NojSLQwSL/4jyHUsPSwQGuzhzEzD2TbTY+8w",
	"ZX7sndEXForhO8qOFF1LoC1vXQVDH4lR95gOig31Ux4GzgAtS5ZddQyBdtRBdZHeSNv3WdwdLODuusF2",
	"YCAw+sWiaiWodsJ+I93aslE8XNtsFGbO2mn1IUMIp2LKFz3sI8qQNlbm2oWrM6D5D7D5xbTF5Uyup5O7",
	"2Q1juHYj7sD1Sb29UTyjh9vakVpugBuinJalFBc0T5x1dYg0pbhwpInNvTH2gVld3IZ39ubw7YkD/3o6",
	"SXOgMqlFhcFVYbvyL7MqWxtg4ID4ompG4fEyuxUlg82vc7ZDi+zlGlwBq0Aa7VXaaKztwVF0FtplPNBm",
	"p73VOQbsErc4CKCs/QON7cq6B9ouAXpBWe6NRh7agaAYXNy4ci1RrhAOcGfXQuAhSu6V3fROd/x0NNS1",
	"gyeFc20psVXYKnKKCN71HxsREm1RSKoFxToZ1iTQZ068KhJz/BKVszRuYOQLZYiDW8eRaUyw8YAwakas",
	"2IAfklcsGMs0UyMU3Q6QwRxRZPqaK0O4WwhX/rfi7PcKCMuAa/NJ4qnsHFQsTOJMzf3r1MgO/bncwNY8",
	"3Qx/FxkjrBHTvfEQiO0CRuim6oF7VKvMfqG1Ocb8ENjjb+DtDmfsXYlbPNWOPhw12xjAddvdFFbr7fM/",
	"Qxi2stvuUsFeeXXFagbmiJb+ZSpZSvEHxPU8VI8jIfe+Kg7DEI8/gM8imUtdFlNbd5oKxs3sg9s9JN2E",
	"Vqi2h36A6nHnA58UViDx5lnK7VbbSpytQK84wYTBmXM7fkMwDuZeQGtOLxc0Vp7FCBkGpsPG+9kyJGtB",
	"fGePe2fzZq5Q0YwEjtS6LbPJaCXIJhumn/h8S4HBTjtaVGgkA6TaUCaYWudXrkRkmIpfUm4Lupp+9ii5",
	"3gqs8cv0uhQSU0lV3OadQcoKmsclhwyx3069zdiK2XKmlYKgXqYbyNaBtlTkao5a/3KDmuMl2ZsGFXnd",
	"bmTsgim2yAFb7NsWC6qQk9eGqLqLWR5wvVbY/OmI5uuKZxIyvVYWsUqQWqhD9ab23CxAXwJwsoft9l+R",
	"r9BnpdgFPDZYdPfz5GD/FRpd7R97sQvA1S3exk0yZCf/dOwkTsfotLNjGMbtRp1FEyNtsflhxrXlNNmu",
	"Y84StnS8bvdZKiinK4iHSRQ7YLJ9cTfRkNbBC89spWSlpdgQpuPzg6aGPw0EcRv2Z8EgqSgKpgvn2VCi",
	"MPTUFMO0k/rhbNllV6rJw+U/ooOw9P6RjhL5sEZTe7/FVo1u3He0gDZap4Ta/OGcNa57X12NHPsqBFi7",
	"qi5ZZXFj5jJLRzEHPflLUkrGNSoWlV4m/yDpmkqaGvY3GwI3Wbx8HqnX1S7Rw28G+IPjXYICeRFHvRwg",
	"ey9DuL7kKy54UhiOkj1ukiaCUznoyYxHi3mO3g0W3D70WKHMjJIMklvVIjcacOo7ER7fMuAdSbFez43o",
	"8cYre3DKrGScPGhldujn92+dlFEIGatJ0xx3J3FI0JLBBQauxTfJjHnHvZD5qF24C/R/rufBi5yBWObP",
	"ckwR+KZiefZLkwTWKXkoKU/XUbv/wnT8talMXS/ZnuNoCZQ15Rzy6HD2zvzV362R2/9fYuw8BeMj23ZL",
	"GdrldhbXAN4G0wPlJzToZTo3E4RYbWfF1FGX+UpkBOdp6m00VNavzhiUK/u9AqVjGQb4wUZ+oH3H6AW2",
	"WhYBnqFUPSPf2Zdl1kBa5QBQmmVFldvUcshWIJ3hsSpzQbMpMeOcvTl8S+ysto+tr2qrda1QmGuvoqPX",
	"B8V5xsUQ+lKp8fjm8eNsD7g0q1Yaq3MoTYsylotmWpz5BpjwFto6UcwLsTMjR1bCVl5+s5MYelgyWRjJ",
	"tB7N8nikCfMfrWm6RtG1xU2GSX58mTlPlSooxl8X1a3r6+C5M3C7SnO20NyUCKNfXDJlHxSBC2inv9W5",
	"oE518ulw7eXJinNLKVEevS1X+TZo98BZh7Y3h0Yh6yD+hoKLrdJ406p7p9grWrCiW8KvV4XfpkDV9WD9",
	"Q1Ep5YKzFMtFBE+Y1CC7x0nG+ApGVNboGqP8EXcnNHK4ooUD63Aih8XBUoKeETrE9Y2VwVezqZY67J8a",
	"X8FYU01WoJXjbJBNff1LZy9hXIGrl4Tv1AR8UsiW/wU5ZNSll9Sm3xuSEcbODwjA35pv75x6hEGl54yj",
	"IOTQ5uJXrUUD307QRnpimqwEKLeedgKh+mD6zDCXLoOrTzP/1gKOYd0XZtnWV9cf6tB77pynzLR9bdra",
	"ygnNz60wRTvpYVm6SYero0blAX3FBxEc8cAk3gQeILcePxxtC7ltdbnjfWoIDS7QYQcl3sM9wqgrhXZK",
	"I1/QvLIUhS2IDXWJJkwzHgHjLePQvAQSuSDS6JWAG4PndaCfSiXVVgQcxdPOgObopYsxNKWdifauQ3U2",
	"GFGCa/RzDG9jU+R0gHHUDRrBjfJN/QCJoe5AmHiNLx85RPZLlqJU5YSoDMOOO0VMY4zDMG5fJrl9AfSP",
	"QV8mst21pPbk3OQmGsokW1TZCnRCsyxWaO4b/ErwK8kqlBzgCtKqLtRVliTFEgztmhR9anMTpYKrqtgy",
	"l29wx+lSEZOj3+EEysdVN4PPCLJfw3qP3py8f/P68OzNkb0vjFpuU8mMzC2hMAzR6LFKgxGdKwXktxCN",
	"v2G/3zoLjoMZFC+OEG1YQNkTIgbULzb4b6yY1jABOZ/6jaO6vAMdO95YvG+P1BPOzdFLFFsl4zGBV9/d",
	"0dFMfbvz2PS/1wOZi1UbkAdOc9/GjMM9irHhN+Z+C7PAexXi7A1YJ2ljDJXw7yCgdlunF7aZJ964vZJx",
	"aLuvS9pvt54MF6ef4h09EEkZJPdTKwZYZ9BQPGU6GP5LtcvC0ZRs5ZRYUT42gg3GsJXs7VuYUUPYUACG",
	"jb8wn3u9xwmwPXUAx96KUB/Z0wfoBx82SErKnKezYRZ9zLoA437I95jQw2aDu4twYbs4SGwl8Qrhw3U2",
	"mtoaeA2UQrGmqmWsdPjIsJIzrP4d1Anpj+V9uheQaiPUB74qCXCTqiFmsuChgy/1NgbUjzr6xpXZ2FZb",
	"o1+/dAez6WUABFkstvbjbHwlicM6IgH9pPjUwAq4e2ugHds7OsJwuYRUs4sdGRf/NFpqE80/9Xqsfcgm",
	"SMBgdcSaf333hup1A9C2hIit8AT1p+4MzlC89TlsHinSooZoMcqp53m3SVRGDCB3SAyJCBXz+FnDm3PC",
	"MFVTBmLBe9htd2hKvgxWAQ/yh245lydJQsOcoi1TXoiY5j5qLtP1Rpl2GHw1lJTRr8M7LAgdYdljVb/g",
	"UD+vG2g15LhfDurSJUpjfkxta/Yp06D8bz4Zzs5in21u6pSjZf+Sysy3iKqqXgtOttxHvUwKX0O2C/Sy",
	"npk18VD92PlIgRGMektzoRhfJUOhg+0QpPDJN3S04nWABY4RriVI9z6B9q9iJ1r4+KltcGxDhXue7DZI",
	"UINFvSxwg6n275taAlgmkdo30Z0TOVyg0VupgU4GGf/Dc25D9mv73QeL+zJ5IzRyR6/JzpR9HwnHVA+J",
	"IdUvibstdweh30brZZzb92pULP2fG1SG1uNSiqxK7QUdHozGxjC2uMYWVhJVGNP+Knuyf46lZt4GKT3n",
	"sJlb+TtdU97U/GkfaytC2TUEKbSd3b5Xg0Bc98lXdgGre4Hzz1Sqp5NSiDwZMBcf96sYdM/AOUvPISPm",
	"7vAxJAOVwMlXaKWs/YGX643P2i9L4JA9nhFi1PKi1BvvGmwX5OxMzh/pbfNf4axZZQuLOH1/9pHHw5+w",
	"5Ie8I3/zw2znagoM87vjVHaQHWUCrgYqKEh6GamLP/atxoizrlurvCEqC0VMSrllzuio893X+SOkHxTr",
	"3q79hCnlPuszFdKajlBa8gadrvDy49Ajhx290T+ZWZ+zYEqFEbWGKShbAEv0haTAKKBe17p2fPP7Kjnm",
	"QwuONSL6qrxC0xjWWQwRYehMXtD84dVxTJQ/RHy411ziCw31uRDJFpXqdt73t3TU3IHudn9T8xM0H/wT",
	"zB5FbZpuKGcXrAuQ+7ppWA+I5iQXzUMEOCS5xDGtEXT/JVm4CMtSQsoU6wSfX/oScLX6ghVRm0d+tutL",
	"u9b5i9B3IGMn8IqSvGvKSWmB/K6BsDn7f3Ig3cDJjVJ5jPp6ZBHB3yD3PdltcgueJfCyjiPmPxlxP9Ys",
	"N1hKbKVhUucORn/esgPbQoSdOAwh4Z7twYED+ob24H666tjl4TpwQysF/XWOvmdbuI1csc3axjoz+sgd",
	"9kHoxRgfRLxomumOThCLEKw4SBBU8tv+b0TCEkuKC/LkCU7w5MnUNf3tafuzYVxPnkQFsAdzf7Re/nXz",
	"xijml6G4PRubNhAi2tmPiuXZLsJoBfw25f0xpPVXFxr9pzww8Ku1hPaPqivNfBPHa3cTEDGRtbYmD6YK",
	"QnlHRPG6brPo28wK0koyvcGMbW84Y79GK+F8V9vana+mzvFzt7wW51Dn/DeW+Up5OeI7Yd9yLow0jG5v",
	"jW89vbmiRZmDOyhfP1r8HZ7943m292z/74t/7L3YS+H5i1d7e/TVc7r/6tk+PP3Hi+d7sL98+WrxNHv6",
	"/Oni+dPnL1+8Sp893188f/nq748MHzIgW0AnPj9o8j/xFY7k8OQ4OTPANjihJaufeDNk7AuA0xRPIhSU",
	"5ZMD/9N/9ydsloqiGd7/OnHpB5O11qU6mM8vLy9nYZf5Ck1xiRZVup77efpPa50c16HRNqUVd9RGvRpS",
	"wE11pHCI396/OT0jhyfHs4ZgJgeTvdnebB8fzimB05JNDibP8Cc8PWvc97kjtsnB5+vpZL4GmqPnyvxR",
	"gJYs9Z/UJV2tQM5cJXTz08XTuRea5p+dGfJ627d5WFRw/rllrc129MS6a/PPPp14e+tWvq6zUgcdRkIx",
	"PKV9PHb+GYX2wd/nqB9Zcpx770K8ZQvgz/qKZdfdHu65xvnn5v3Ua3tec4j5EmxQPQ2eW50SpgldCIkp",
	"tTpdmyPqc/mYaj+3W9PbcWbozPR6Xb8lG5QxOvjQUy3sQMSPhIfSUFxzZlozNWxRywrCyjo102+1b1j/",
	"h73k1afP+9P9veu/Gdbu/nzx7HqkU/B18xTtac23Rzb8hIlwaN7Eo/R0b+8OLy0d8vBdXNyk4EGv6PPY",
	"VZkUQyYDt1WdgUiNjB0JO53hBx7jfH7DFW81wrRCcCIPJ3xDM+LzTHDu/Yeb+5ijS9awWGKvkOvp5MVD",
	"rv6YG5KnOcGWQQZ2f+t/5udcXHLf0tz3VVFQufHHWLWYgn8hGm8VulJokpPsgmqYfEKbbywgcYC5KE1v",
	"wVxOTa8vzOWhmAtu0n0wl/ZA98xcnt7wgP/1V/yFnf7V2OmpZXfj2akT5Wwq49w+7thIeL1C/yuI5lRi",
	"diPd9mpzl8N+B7r3CPXkjizmT3uP+v/vc/J87/nDQdCuUv0DbMg7ocm3aEz+i57ZccdnmyTU0YyyrEfk",
	"lv2D0t+IbLMFQ4ValS79KCKXLBg3IPdvl/6zh71Hos9hQ2xskPcBc5FBTx66viMP+Mu+Z/2Fh3zhIdJO",
	"/+zhpj8FecFSIGdQlEJSyfIN+ZnXyeO3V+uyLBp33T76PZ5mtJFUZLACnjiGlSxEtvGFA1sDnoM1YvcE",
	"lfnndvVvaygbNEsd4e/1o4R9oBcbcnzUk2Bsty6n/WaDTTsaY0Qn7IK4VTPs8qIBZWwbmZuFrIQmFguZ",
	"W9QXxvOF8dxJeBl9eGLyS1Sb8Iac7p089VVUYnWGqO5PPUbn+FOP671sdF+fiekvNj4dMhJ8sIlUXTR/",
	"YQlfWMLdWMJ3EDmMeGodk4gQ3W0svX0GgaG4WfcNHQx08M2rnEqiYKyZ4hBHdMaJh+ASD62kRXFldTTK",
	"CVwxG7UZ2bD71du+sLgvLO4v5LXazWjagsiNNZ1z2BS0rPUbta50Ji5t9cEoV8TC/DR3VXwxorSO2dCC",
	"+AGajFfyk6sWkG8wjJZlRozTrAAjUtW8znT2eQxNgLcZoXlMecU4ToCsAmex5appkEumIBXcPj3a8bU5",
	"yN5ZnTDGZH+vADmaw42DcTJtOVvcNkaKQ99Z/ur7Rq632NKRKmzsez8eo35ctPX3/JIynSyFdHmmiL5+",
	"Zw00n7siXJ1fm4ISvS9YJSP4MQjsiP86rx9TiH7sBq/EvrqIEd+oiU4Lo71wg+s4rw+fzD5hLV63903w",
	"0sF8jslZa6H0fHI9/dwJbAo/fqq35nN9Lbstuv50/X8DAAD//89W3bJFvAAA",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}

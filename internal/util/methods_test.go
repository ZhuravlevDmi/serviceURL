package util

//func TestCheckMapUrl(t *testing.T) {
//	type args struct {
//		mapURL map[string]string
//		path   string
//	}
//	tests := []struct {
//		name string
//		args args
//		want string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := CheckMapUrl(tt.args.mapURL, tt.args.path); got != tt.want {
//				t.Errorf("CheckMapUrl() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestGenerateMiniUrl(t *testing.T) {
//	tests := []struct {
//		name string
//		want string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := GenerateMiniUrl(); got != tt.want {
//				t.Errorf("GenerateMiniUrl() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestSetMapUrl(t *testing.T) {
//	type args struct {
//		mapURL map[string]string
//		path   string
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    string
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := SetMapUrl(tt.args.mapURL, tt.args.path)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("SetMapUrl() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("SetMapUrl() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

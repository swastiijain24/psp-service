package services

type VpaService struct {

}

func NewVpaService() *VpaService{
	return &VpaService{}
}

func (s* VpaService) ResolveVpa(VPA string) (string){
	//this will either resolve by itself or call the psp to resolve the add 
	return ""
}
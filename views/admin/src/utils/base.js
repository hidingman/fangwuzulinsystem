const base = {
    get() {
        return {
            url : "http://localhost:8080/fangwuzulinsystem/",
            name: "fangwuzulinsystem",
            // 退出到首页链接
            indexUrl: 'http://localhost:8080/fangwuzulinsystem/front/h5/index.html'
        };
    },
    getProjectName(){
        return {
            projectName: "骄阳房屋租赁公司业务管理系统"
        } 
    }
}
export default base

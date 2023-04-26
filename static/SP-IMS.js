var v1 = new Vue({
  el: "#app1",
  methods: {
    handleClick(row) {
      console.log(row);
    },



    async getDate(pageNum) {
      const { data: res } = await axios({
        method: "get",
        url: "http://localhost:18080/program/all",
        params: {
          pageNum,
        },
        headers: {
          "content-type": "application/json",
        },
      });
      console.log(res.data.programResults)
      // res.data.forEach((item) => {
      //   item.actorList = this.changeArrToStr(item.actorList);
      // });

      this.tableData = res.data.programResults;
      console.log(this.tableData)
    },
    changeArrToStr(arr) {
      let newArr = [];
      arr.forEach((item) => newArr.push(item.name));
      return newArr.join(",");
    },
    pageChange(pageNum) {
      console.log(pageNum);
      if (pageNum <= 0) {
        return;
      } else {
        this.getDate(pageNum);
      }
    },
    addConfirm(){
      this.isAdd=false
      this.addData(this.addForm)
    },
    add(){
      this.isAdd=true
    },
    async addData(addForm) {
      const { data: res } = await axios({
        method: "post",
        url: "http://localhost:18080/program/add",
        data: {
          type:addForm.category,
          name:addForm.program,
          actors:addForm.actor,
          point:addForm.point
        },
        headers: {
          "content-type": "application/json",
        },
      });
      console.log(res)
      if(res.code=='00000'){
        this.getDate(this.pageNum);
        this.$message({
          message: '新增成功',
          type: 'success'
        });
      } else {
        this.$message.error(res.message);
      }
      
    },
    async search(searchForm) {
      const { data: res } = await axios({
        method: "get",
        url: "http://localhost:18080/program/search",
        params: {
          type:searchForm.category,
          num:searchForm.actorNum,
          name:searchForm.program
        },
        headers: {
          "content-type": "application/json",
        },
      });
      console.log(res)
      // res.data.forEach((item) => {
      //   item.actorList = this.changeArrToStr(item.actorList);
      // });

      this.tableData = res.data.programResults;
      console.log(res.data.programResults)
      console.log(this.tableData)
    },
    edit(row){
      console.log(row,"@@@@")
      this.editForm.index = row.id;
      this.editForm.category=row.typeName
      this.editForm.program=row.name
      this.editForm.actor=row.actorList
      this.editForm.point=row.view
      this.isEdit=true
      
    },
    editConfirm(){
      this.isEdit=false,
      this.editData(this.editForm)
    },
    async editData(editForm) {
      const { data: res } = await axios({
        method: "post",
        url: "http://localhost:18080/program/update",
        data: {
          id:editForm.index,
          type:editForm.category,
          name:editForm.program,
          actors:editForm.actor,
          point:editForm.point
        },
        headers: {
          "content-type": "application/json",
        },
      });
      console.log(editForm)
      if(res.code=='00000'){
        this.getDate(this.pageNum);
        this.$message({
          message: '编辑成功',
          type: 'success'
        });
      } else {
        this.$message.error(res.message);
      }
      
    },
    async deleteData(index) {
      const { data: res } = await axios({
        method: "post",
        url: "http://localhost:18080/program/deleted",
        data: {
          id:index
        },
        headers: {
          "content-type": "application/json",
        },
      });
      console.log(res)
      if(res.code=='00000'){
        this.getDate(this.pageNum);
        this.$message({
          message: '删除成功',
          type: 'success'
        });
      } else {
        this.$message.error(res.message);
      }
      
    },
  },
  data() {
    return {
      cateOptions: [{
        value: 1,
        label: '文学'
      }, {
        value: 2,
        label: '科技'
      }, {
        value: 3,
        label: '艺术'
      }, {
        value: 4,
        label: '历史'
      }, {
        value: 5,
        label: '语言'
      }, {
        value: 6,
        label: '自然科学'
      }],
      
      addForm:{
        category:null,
        program:null,
        actor:null,
        point:null
      },
      searchForm:{
        category:null,
        actorNum:null,
        program:null,
      },
      editForm:{
        index:null,
        category:null,
        program:null,
        actor:null,
        point:null
      },
      value: '',
      isEdit:false,
      isAdd:false,
      tableData: [],
      pageNum: 1,
    };
  },
  created() {
    this.getDate(this.pageNum);
  },
});






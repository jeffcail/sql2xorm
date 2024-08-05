new Vue({
    el: '#app',
    data: {
        sqlInput: '',
        structOutput: ''
    },
    methods: {
        async convertSQL() {
            if (!this.sqlInput.trim()) return;  // 不发送空请求
            try {
                const response = await axios.post('/gen', {
                    sql: this.sqlInput
                });
                this.structOutput = response.data.struct;
            } catch (error) {
                console.error("There was an error!", error);
                this.structOutput = error;
            }
        }
    }
});

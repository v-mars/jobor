<template>
  <div>
    <el-row :gutter="15" class="system_state">
      <el-col :span="12">
        <el-card v-if="state.os" class="card_item">
          <div slot="header">Runtime</div>
          <div>
            <el-row :gutter="10">
              <el-col :span="12">os:</el-col>
              <el-col :span="12" v-text="state.os.goos"></el-col>
            </el-row>
            <el-row :gutter="10">
              <el-col :span="12">cpu nums:</el-col>
              <el-col :span="12" v-text="state.os.numCpu"></el-col>
            </el-row>
            <el-row :gutter="10">
              <el-col :span="12">compiler:</el-col>
              <el-col :span="12" v-text="state.os.compiler"></el-col>
            </el-row>
            <el-row :gutter="10">
              <el-col :span="12">go version:</el-col>
              <el-col :span="12" v-text="state.os.goVersion"></el-col>
            </el-row>
            <el-row :gutter="10">
              <el-col :span="12">goroutine nums:</el-col>
              <el-col :span="12" v-text="state.os.numGoroutine"></el-col>
            </el-row>
          </div>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card v-if="state.disk" class="card_item">
          <div slot="header">Disk</div>
          <div>
            <el-row :gutter="10">
              <el-col :span="12">
                <el-row :gutter="10">
                  <el-col :span="12">total (MB)</el-col>
                  <el-col :span="12" v-text="state.disk.totalMb"></el-col>
                </el-row>
                <el-row :gutter="10">
                  <el-col :span="12">used (MB)</el-col>
                  <el-col :span="12" v-text="state.disk.usedMb"></el-col>
                </el-row>
                <el-row :gutter="10">
                  <el-col :span="12">total (GB)</el-col>
                  <el-col :span="12" v-text="state.disk.totalGb"></el-col>
                </el-row>
                <el-row :gutter="10">
                  <el-col :span="12">used (GB)</el-col>
                  <el-col :span="12" v-text="state.disk.usedGb"></el-col>
                </el-row>
              </el-col>
              <el-col :span="12">
                <el-progress
                  type="dashboard"
                  :percentage="state.disk.usedPercent"
                  :color="colors"
                ></el-progress>
              </el-col>
            </el-row>
          </div>
        </el-card>
      </el-col>
    </el-row>
    <el-row :gutter="15" class="system_state">
      <el-col :span="12">
        <el-card
          class="card_item"
          v-if="state.cpu"
          :body-style="{ height: '180px', 'overflow-y': 'scroll' }"
        >
          <div slot="header">CPU</div>
          <div>
            <el-row :gutter="10">
              <el-col :span="12">physical number of cores:</el-col>
              <el-col :span="12" v-text="state.cpu.cores"> </el-col>
            </el-row>
            <template v-for="(item, index) in state.cpu.cpus">
              <el-row :key="index" :gutter="10">
                <el-col :span="12">core {{ index }}:</el-col>
                <el-col :span="12"
                ><el-progress
                  type="line"
                  :percentage="+item.toFixed(0)"
                  :color="colors"
                ></el-progress
                ></el-col>
              </el-row>
            </template>
          </div>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card class="card_item" v-if="state.ram">
          <div slot="header">Ram</div>
          <div>
            <el-row :gutter="10">
              <el-col :span="12">
                <el-row :gutter="10">
                  <el-col :span="12">total (MB)</el-col>
                  <el-col :span="12" v-text="state.ram.totalMb"></el-col>
                </el-row>
                <el-row :gutter="10">
                  <el-col :span="12">used (MB)</el-col>
                  <el-col :span="12" v-text="state.ram.usedMb"></el-col>
                </el-row>
                <el-row :gutter="10">
                  <el-col :span="12">total (GB)</el-col>
                  <el-col :span="12" v-text="state.ram.totalMb / 1024"></el-col>
                </el-row>
                <el-row :gutter="10">
                  <el-col :span="12">used (GB)</el-col>
                  <el-col
                    :span="12"
                    v-text="(state.ram.usedMb / 1024).toFixed(2)"
                  ></el-col>
                </el-row>
              </el-col>
              <el-col :span="12">
                <el-progress
                  type="dashboard"
                  :percentage="state.ram.usedPercent"
                  :color="colors"
                ></el-progress>
              </el-col>
            </el-row>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>

  export default {
    name: "State",
    data() {
      return {
        timer:null,
        state: {},
        colors: [
          { color: "#5cb87a", percentage: 20 },
          { color: "#e6a23c", percentage: 40 },
          { color: "#f56c6c", percentage: 80 },
        ],
        loading: false,
        url: this.$store.state.urls.sys_state_url,
      };
    },
    created() {
      this.getData();
      this.timer = setInterval(() => {
        this.getData();
      }, 1000*10);
    },
    beforeDestroy(){
      clearInterval(this.timer)
      this.timer = null
    },
    methods: {
      getData: async function () {
        // let data = {"pageSize": this.limit, "pageNumber": this.page, name: this.SearchForm.name}
        this.loading = true
        try {
          let url = this.url
          let response = await this.$store.dispatch("common/Query", {url: url, data: {}})
          this.state = response.data.data
        } catch (e) {

        } finally {
          this.loading = false
        }
      },

      async reload() {
        const { data } = await getSystemState();
        this.state = data.server;
      },
    },
  };
</script>

<style lang="scss" scoped>
  /deep/ .system_state {
    padding: 10px;
  }

  /deep/ .card_item {
    height: 280px;
  }
</style>

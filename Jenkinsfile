pipeline {
  agent any

  options {
      timeout(time: 5, unit: 'MINUTES')
      // 不允许同时执行流水线, 防止同时访问共享资源等
      disableConcurrentBuilds()
      // 显示具体的构建流程时间戳
      timestamps()
  }

  environment {
    PROJECT_NAME = 'advproject'
    DOCKER_REGISTER = 'ccr.ccs.tencentyun.com'
    DOCKER_NAMESPCAE = 'darkreunion'
    CHANGE_LOG = sh returnStdout: true, script: 'git log --pretty=format:\'%h - %an,%ar : %s\' --since=\'1 hours\' | head -n 1'
  }

  stages {
    stage('Test') {
      agent {
        docker {
          image 'golang:1.13'
        }
      }
      steps {
        sh '''go env -w GOPROXY=https://goproxy.cn,direct
        go test -v ./test'''
      }
    }

    stage('Build Docker') {
      steps {
        script {
          withDockerRegistry(credentialsId: 'tcloud-docker-reg', url: 'https://ccr.ccs.tencentyun.com') {
            def image = docker.build("${env.DOCKER_REGISTER}/${env.DOCKER_NAMESPCAE}/${env.PROJECT_NAME}:build-${BUILD_NUMBER}")
            image.push()
          }
        }
      }
      post {
        success {
          sh 'docker rmi `docker images | awk \'/^<none>/ { print $3 }\'`'
          sh "docker rmi ${env.DOCKER_REGISTER}/${env.DOCKER_NAMESPCAE}/${env.PROJECT_NAME}:build-${BUILD_NUMBER}"
        }
      }
    }

    stage('Deploy'){
      steps {
        withKubeConfig([credentialsId: 'kubctl-config', serverUrl: "${env.K8S_SERVER}"]) {
          sh 'kubectl get pods'
        }
      }
    }
  }

  post {
    always {
      withCredentials([string(credentialsId: 'PUSH_KEY', variable: 'PUSH_KEY')]) {
        sh "curl -s -d 'text=项目 ${currentBuild.projectName} 集成结果: ${currentBuild.result}' -d 'desp=change log: `${env.CHANGE_LOG}`' 'https://sc.ftqq.com/${PUSH_KEY}.send'"
      }
    }
  }
}
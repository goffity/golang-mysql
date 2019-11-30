node{
    stage("checkout"{
        checkout scm
    })
    stage("build"){
        sh "go build"
    }
    stage("deploy"){
        // sshPublisher(
        //     publishers: [
        //         sshPublisherDesc(
        //             configName: 'server-1', transfers: [
        //                 sshTransfer(
        //                     cleanRemote: false, excludes: '', execCommand: './go.sh', execTimeout: 120000, flatten: false, makeEmptyDirs: false, noDefaultExcludes: false, patternSeparator: '[, ]+', remoteDirectory: '', remoteDirectorySDF: false, removePrefix: '', sourceFiles: 'dist/*'
        //                 )
        //             ], usePromotionTimestamp: false, useWorkspaceInPromotion: false, verbose: false
        //         )
        //     ]
        // )
        sh "scp README.md vagrant@192.168.31.241:/usr/share/nginx/html/README.md"
    }
}
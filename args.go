// Code generated by gen.sh. DO NOT EDIT.
package ctags

var ctagsArgs = []string{
	`--langmap=java:+.apex.cls.trigger`,
	`--regex-java=/^[ \t]*(private|protected|global|public|webservice)?[ \t]*(static)?[ \t]*[a-zA-Z0-9_\.]+[ \t]*(<[ \t]*[a-zA-Z0-9_\.]+[ \t]*,[ \t]*[a-zA-Z0-9_\.]+[ \t]*>[ \t]*)?([a-zA-Z0-9_]+)[ \t]*[{][ \t]*(get|set)/\4/P,properties/i`,
	`--regex-java=/^[ \t]*(private|protected|global|public)?[ \t]*((with|without) sharing)?[ \t]*(abstract|virtual)?[ \t]*class[ \t]+([a-zA-Z0-9_]+)/\5/classes/i`,
	`--map-C++=+.pc`,
	`--map-C++=+.pcc`,
	`--regex-clojure=/def ([A-Za-z0-9_!?+*<>=-]+)/\1/v,variable/`,
	`--regex-clojure=/defmacro ([A-Za-z0-9_!?+*<>=-]+)/\1/m,macro/`,
	`--regex-clojure=/defprotocol ([A-Za-z0-9_!?+*<>=-]+)/\1/p,protocol/`,
	`--regex-clojure=/defrecord ([A-Za-z0-9_!?+*<>=-]+)/\1/r,record/`,
	`--regex-clojure=/deftype ([A-Za-z0-9_!?+*<>=-]+)/\1/t,type/`,
	`--regex-clojure=/s\/def ([A-Za-z0-9_!?+*<>=-]+)/\1/v,variable/`,
	`--regex-clojure=/s\/defn ([A-Za-z0-9_!?+*<>=-]+)/\1/f,function/`,
	`--regex-clojure=/s\/defmacro ([A-Za-z0-9_!?+*<>=-]+)/\1/m,macro/`,
	`--regex-clojure=/s\/defprotocol ([A-Za-z0-9_!?+*<>=-]+)/\1/p,protocol/`,
	`--regex-clojure=/s\/defrecord ([A-Za-z0-9_!?+*<>=-]+)/\1/r,record/`,
	`--regex-clojure=/s\/deftype ([A-Za-z0-9_!?+*<>=-]+)/\1/t,type/`,
	`--map-CSS=+.scss`,
	`--map-CSS=+.less`,
	`--map-CSS=+.sass`,
	`--langdef=GraphQL`,
	`--langmap=GraphQL:.graphql`,
	`--regex-graphql=/^[ \t]*([_A-Za-z][_A-Za-z0-9]*)[\(:]/\1/v,field/`,
	`--regex-graphql=/^[ \t]*enum[ \t]+([_A-Za-z][_0-9A-Za-z]*)/\1/e,enum/`,
	`--regex-graphql=/^[ \t]*fragment[ \t]+([_A-Za-z][_0-9A-Za-z]*)/\1/f,fragment/`,
	`--regex-graphql=/^[ \t]*input[ \t]+([_A-Za-z][_0-9A-Za-z]*)/\1/i,input/`,
	`--regex-graphql=/^[ \t]*interface[ \t]+([_A-Za-z][_0-9A-Za-z]*)/\1/n,interface/`,
	`--regex-graphql=/^[ \t]*mutation[ \t]+([_A-Za-z][_0-9A-Za-z]*)/\1/m,mutation/`,
	`--regex-graphql=/^[ \t]*query[ \t]+([_A-Za-z][_0-9A-Za-z]*)/\1/q,query/`,
	`--regex-graphql=/^[ \t]*type[ \t]+([_A-Za-z][_0-9A-Za-z]*)/\1/t,type/`,
	`--regex-graphql=/^[ \t]*scalar[ \t]+([_A-Za-z][_0-9A-Za-z]*)/\1/t,type/`,
	`--langdef=Groovy`,
	`--langmap=Groovy:.groovy`,
	`--regex-groovy=/^[ \t]*package[ \t]+([a-zA-Z0-9.-_]+)/\1/p,package/`,
	`--regex-groovy=/^[ \t]*((private|public)[ \t]*)?(abstract|final|static)?[ \t]*class[ \t]+([A-Za-z0-9_]+)/\4/c,class/`,
	`--regex-groovy=/^[ \t]*((private|public)[ \t]*)?interface[ \t]+([A-Za-z0-9_]+)/\3/n,interface/`,
	`--regex-groovy=/^[ \t]*((private|public)[ \t]*)?trait[ \t]+([A-Za-z0-9_]+)/\3/t,trait/`,
	`--regex-groovy=/^[ \t]*((private|public)[ \t]*)?enum[ \t]+([A-Za-z0-9_]+)/\3/e,enum/`,
	`--regex-groovy=/^[ \t]*(public|private|protected[ \t]*)*([a-zA-Z0-9_]{3,})\([A-Za-z0-9 _,]*\)[ \t]+/\2/r,constructor/`,
	`--regex-groovy=/^[ \t]*((abstract|final|public|private|protected|static)[ \t]*)*(def|void|byte|int|short|long|float|double|boolean|char|[A-Z][a-zA-Z0-9_]*)[ \t]+([a-zA-Z0-9_]+)\(/\4/m,method/`,
	`--regex-groovy=/^[ \t]*((final|public|private|protected|static|synchronized)[ \t]*)*(def|byte|int|short|long|float|double|boolean|char|[A-Z][A-Za-z0-9_]*)[ \t]+([a-zA-Z0-9_]+)[ \t]*([\/]+.*)?$/\4/v,field/`,
	`--langdef=Jsonnet`,
	`--langmap=Jsonnet:.jsonnet`,
	`--langmap=Jsonnet:.libsonnet`,
	`--regex-jsonnet=/^[ \t]*local ([A-Za-z0-9_]+)/\1/v,variable/`,
	`--regex-jsonnet=/^[ \t]*([A-Za-z0-9_]+):::?/\1/m,member/`,
	`--regex-jsonnet=/^[ \t]*([A-Za-z0-9_]+)\([^)]*\):::?/\1/f,function/`,
	`--langdef=scala`,
	`--langmap=scala:.scala`,
	`--regex-scala=/^[ \t]*((abstract|final|sealed|implicit|lazy)[ \t]*)*(private[^ ]*|protected)?[ \t]*class[ \t]+([a-zA-Z0-9_]+)/\4/c,class/`,
	`--regex-scala=/^[ \t]*((abstract|final|sealed|implicit|lazy)[ \t]*)*(private[^ ]*|protected)?[ \t]*object[ \t]+([a-zA-Z0-9_]+)/\4/o,object/`,
	`--regex-scala=/^[ \t]*((abstract|final|sealed|implicit|lazy)[ \t]*)*(private[^ ]*|protected)?[ \t]*((abstract|final|sealed|implicit|lazy)[ \t ]*)*case class[ \t ]+([a-zA-Z0-9_]+)/\6/c,class/`,
	`--regex-scala=/^[ \t]*((abstract|final|sealed|implicit|lazy)[ \t]*)*(private[^ ]*|protected)?[ \t]*case object[ \t]+([a-zA-Z0-9_]+)/\4/o,object/`,
	`--regex-scala=/^[ \t]*((abstract|final|sealed|implicit|lazy)[ \t]*)*(private[^ ]*|protected)?[ \t]*trait[ \t]+([a-zA-Z0-9_]+)/\4/i,interface/`,
	`--regex-scala=/^[ \t]*type[ \t]+([a-zA-Z0-9_]+)/\1/T,type/`,
	`--regex-scala=/^[ \t]*((abstract|final|sealed|implicit|lazy|override|private[^ ]*(\[[a-z]*\])*|protected)[ \t]*)*def[ \t]+([a-zA-Z0-9_]+)/\4/m,method/`,
	`--regex-scala=/^[ \t]*((abstract|final|sealed|implicit|lazy|override|private[^ ]*|protected)[ \t]*)*val[ \t]+([a-zA-Z0-9_]+)/\3/v,variable/`,
	`--regex-scala=/^[ \t]*((abstract|final|sealed|implicit|lazy|override|private[^ ]*|protected)[ \t]*)*var[ \t]+([a-zA-Z0-9_]+)/\3/v,variable/`,
	`--regex-scala=/^[ \t]*package[ \t]+([a-zA-Z0-9_.]+)/\1/p,package/`,
	`--langdef=swift`,
	`--langmap=swift:.swift`,
	`--regex-swift=/^[[:space:]]*((associatedtype|class|deinit|enum|extension|fileprivate|func|import|init|inout|internal|let|open|operator|private|protocol|public|static|struct|subscript|typealias|var|final)[[:space:]])*class[[:space:]]+([[:alnum:]_]+)/\3/c,class/`,
	`--regex-swift=/^[[:space:]]*public[[:space:]]*((associatedtype|class|deinit|enum|extension|fileprivate|func|import|init|inout|internal|let|open|operator|private|protocol|public|static|struct|subscript|typealias|final)[[:space:]])*let[[:space:]]+([[:alnum:]_]+)/\3/C,constant/`,
	`--regex-swift=/^[[:space:]]*public[[:space:]]*((associatedtype|class|deinit|enum|extension|fileprivate|func|import|init|inout|internal|let|open|operator|private|protocol|public|static|struct|subscript|typealias|final|lazy|mutating|nonmutating|optional|override|required|unowned|weak)[[:space:]])*var[[:space:]]+([[:alnum:]_]+)/\3/v,variable/`,
	`--regex-swift=/^[[:space:]]*((associatedtype|class|deinit|enum|extension|fileprivate|func|import|init|inout|internal|let|open|operator|private|protocol|public|static|struct|subscript|typealias|var)[[:space:]])*enum[[:space:]]+([[:alnum:]_]+)/\3/e,enum/`,
	`--regex-swift=/^[[:space:]]*((associatedtype|class|deinit|enum|extension|fileprivate|func|import|init|inout|internal|let|open|operator|private|protocol|public|static|struct|subscript|typealias|var|convenience|dynamic|final|mutating|nonmutating|optional|override|required)[[:space:]])*func[[:space:]]+([[:alnum:]_]+)/\3/f,function/`,
	`--regex-swift=/^[[:space:]]*((associatedtype|class|deinit|enum|extension|fileprivate|func|import|init|inout|internal|let|open|operator|private|protocol|public|static|struct|subscript|typealias|var)[[:space:]])*protocol[[:space:]]+([[:alnum:]_]+)/\3/i,interface/`,
	`--regex-swift=/^[[:space:]]*((associatedtype|class|deinit|enum|extension|fileprivate|func|import|init|inout|internal|let|open|operator|private|protocol|public|static|struct|subscript|typealias|var)[[:space:]])*struct[[:space:]]+([[:alnum:]_]+)/\3/s,struct/`,
	`--regex-swift=/^[[:space:]]*((associatedtype|class|deinit|enum|extension|fileprivate|func|import|init|inout|internal|let|open|operator|private|protocol|public|static|struct|subscript|typealias|var)[[:space:]])*extension[[:space:]]+([[:alnum:]_]+)/\3/d,define/`,
	`--regex-swift=/^[[:space:]]*((associatedtype|class|deinit|enum|extension|fileprivate|func|import|init|inout|internal|let|open|operator|private|protocol|public|static|struct|subscript|typealias|var)[[:space:]])*typealias[[:space:]]+([[:alnum:]_]+)/\3/a,alias/`,
	`--langdef=tsx`,
	`--langmap=tsx:.tsx`,
	`--regex-tsx=/^[ \t]*(export[ \t]+([a-z]+[ \t]+)?)?class[ \t]+([a-zA-Z0-9_$]+)/\3/c,class/`,
	`--regex-tsx=/^[ \t]*(declare[ \t]+)?namespace[ \t]+([a-zA-Z0-9_$]+)/\2/n,module/`,
	`--regex-tsx=/^[ \t]*(export[ \t]+)?module[ \t]+([a-zA-Z0-9_$]+)/\2/n,module/`,
	`--regex-tsx=/^[ \t]*(export[ \t]+)?(default[ \t]+)?(async[ \t]+)?function[ \t]+([a-zA-Z0-9_$]+)/\4/f,function/`,
	`--regex-tsx=/^[ \t]*export[ \t]+(var|let|const)[ \t]+([a-zA-Z0-9_$]+)/\2/v,variable/`,
	`--regex-tsx=/^[ \t]*(var|let|const)[ \t]+([a-zA-Z0-9_$]+)[ \t]*=[ \t]*function[ \t]*[*]?[ \t]*\(\)/\2/v,variable/`,
	`--regex-tsx=/^[ \t]*(export[ \t]+)?(public|protected|private)[ \t]+(static[ \t]+)?(abstract[ \t]+)?(((get|set)[ \t]+)|(async[ \t]+[*]*[ \t]*))?([a-zA-Z1-9_$]+)/\9/m,member/`,
	`--regex-tsx=/^[ \t]*(export[ \t]+)?interface[ \t]+([a-zA-Z0-9_$]+)/\2/i,interface/`,
	`--regex-tsx=/^[ \t]*(export[ \t]+)?type[ \t]+([a-zA-Z0-9_$]+)/\2/t,type/`,
	`--regex-tsx=/^[ \t]*(export[ \t]+)?enum[ \t]+([a-zA-Z0-9_$]+)/\2/e,enum/`,
}

package _1166

type FileSystem struct {
	level map[int][]int
	inode map[string]int
	data  map[int]int
}

func Constructor() FileSystem {
	level := map[int][]int{}
	inode := map[string]int{}
	data := map[int]int{}

	inode["/"] = 0
	level[0] = []int{}
	fs := FileSystem{level: level, inode: inode, data: data}

	return fs
}

func getParent(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			if i == 0 {
				return "/"
			}
			return path[:i]
		}
	}
	return path
}

func (this *FileSystem) CreatePath(path string, value int) bool {
	if path == "" || path == "/" {
		return false
	}

	//路径存在
	if _, ok := this.inode[path]; ok {
		return false
	}

	parent := getParent(path)
	if parentInode, ok := this.inode[parent]; !ok {
		//父目录不存在
		return false
	} else {
		newInode := len(this.inode)
		this.inode[path] = newInode
		this.data[newInode] = value
		this.level[parentInode] = append(this.level[parentInode], newInode)
		return true
	}
}

func (this *FileSystem) Get(path string) int {
	if inode, ok := this.inode[path]; ok {
		return this.data[inode]
	} else {
		return -1
	}
}

/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.CreatePath(path,value);
 * param_2 := obj.Get(path);
 */

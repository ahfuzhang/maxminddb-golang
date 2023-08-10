跟进整个执行流程，搞清楚 decode 的各个细节

* 先解析 ip
func (r *Reader) lookupPointer(ip net.IP) (uint, int, net.IP, error)
* 通过 ip 地址，逐个 bit 在树里面查找
func (r *Reader) traverseTree(ip net.IP, node, bitCount uint) (uint, int)


offset, err := r.resolveDataPointer(pointer)
  offset=59745

func (d *decoder) decode(offset uint, result reflect.Value, depth int) (uint, error) {
  offset = 79754

typeNum, size, newOffset, err := d.decodeCtrlData(offset)
  typeNum _Map
  size = 3
  newOffset = 59755

func (d *decoder) decodeFromType(
	dtype dataType,
	size uint,
	offset uint,
	result reflect.Value,
	depth int,
) (uint, error)

func (d *decoder) unmarshalMap(
	size uint,
	offset uint,
	result reflect.Value,
	depth int,
) (uint, error)
  typeNum _Map
  size = 3
  newOffset = 59755
  depth = 1

func (d *decoder) decodeStruct(
	size uint,
	offset uint,
	result reflect.Value,
	depth int,
) (uint, error)

key, offset, err = d.decodeKey(offset)  // 解析 map 的逐个 key
  key = "continent"
  offset = 59757

func (d *decoder) nextValueOffset(offset, numberToSkip uint)
    // 找不到需要的字段，跳到下一个字段
    offset = 59757
    numberToSkip = 1

d.decodeCtrlData(offset)  // 再次回到解析头
   typeNum = _Pointer
   size = 8
   offset = 59758

_, offset, err = d.decodePointer(size, offset)  // 上面发现是 pointer 类型，因此解析 pointer
   offset = 59760

return d.nextValueOffset(offset, numberToSkip-1)

==============
map 的第二个 key:
"country"

 是需要的字段，然后调用 decode:
   offset = 59762
   depth = 1
offset, err = d.decode(offset, result.Field(j), depth)

typeNum, size, newOffset, err := d.decodeCtrlData(offset)
    // 解析当前 map value 的控制头
   typeNum = _Pointer
   size = 8
   offset = 59763

// value 是 pointer， 解析 pointer 的值
func (d *decoder) decodeFromType(
	dtype dataType,
	size uint,
	offset uint,
	result reflect.Value,
	depth int,
) (uint, error)   

// 转到具体的解析 pointer 的代码去
func (d *decoder) unmarshalPointer(
	size, offset uint,
	result reflect.Value,
	depth int,
) (uint, error)
    newOffset = 59765
    pointer = 8360

d.decodePointer()
d.decode(pointer, result, depth)

// pointer 类型的值
func (d *decoder) decode(offset uint, result reflect.Value, depth int)
   offset = 8360

// 第二层，又是一个 map
typeNum, size, newOffset, err := d.decodeCtrlData(offset)
   typeNum = _Map
   size = 3
   newOffset = 8361

decodeFromType()
d.unmarshalMap(size, offset, result, depth)
func (d *decoder) decodeStruct()

key 0: "geoname_id"
   d.nextValueOffset(offset, 1)  // 跳过
key 1: "iso_code"
      offset 8369
   d.decode(offset, result.Field(j), depth)  // 需要的字段，再进入 decode
     type = _String
     size = 2
     newOffset = 8370


d.unmarshalString(size, offset, result)

func (d *decoder) decodeString(size, offset uint) 
   newOffset: 8372
   value: CL
   "45.232.32.87"


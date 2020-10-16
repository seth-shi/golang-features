var list = [
    {
        title: '江哲',
        description: '须交有道之人，莫结无义之友。饮清静之茶，莫贪花色之酒。开方便之门，闲是非之口。',
        sex: 'boys'
    },
    {
        title: '欧阳锋',
        description: '“我欲”是贫穷的标志。事能常足，心常惬，人到无求品自高。',
        sex: 'boys'
    },
    {
        title: '江小白',
        description: '势不可使尽，福不可享尽，便宜不可占尽，聪明不可用尽。',
        sex: 'boys'
    }
]



createList(list);
function createList(list) {
    var str = '';
    list.forEach(function (ele, index) {
        str += '<li>\
            <div class="tit">\
            <h1 class="title">'+ ele.title + '</h1>\
            <p class="description">'+ ele.description + '</p>\
            </div>\
            </li>'
    })

    $('.container').html(str);
}

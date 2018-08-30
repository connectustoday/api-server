import { prop, Typegoose, ModelType, InstanceType } from 'typegoose';

class IComment {

	@prop({ required: true })
        schema_version: number;

	@prop({ required: true })
	id: number;

	@prop({ required: true })
	account: string;

	@prop({ required: true })
	content: string;

	@prop({ required: true })
	created_at: number;

	@prop()
	multimedia?: IAttachment;

	@prop({ required: true })
	likes_count: number;

	@prop({ required: true })
	comments_count: number;

	@prop()
	likes?: Array<string>;

	@prop()
	comments?: Array<IComment>;
}
